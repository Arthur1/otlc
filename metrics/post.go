package metrics

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkMetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"

	_ "google.golang.org/grpc/encoding/gzip"
)

type Poster struct {
	endpoint string
	headers  map[string]string
}

func NewPoster(endpoint string, headers map[string]string) *Poster {
	return &Poster{
		endpoint: endpoint,
		headers:  headers,
	}
}

type PostParams struct {
	Name           string
	Description    string
	ResourceAttrs  map[string]string
	ScopeAttrs     map[string]string
	DataPointAttrs map[string]string
	DataPointValue float64
}

func (p *Poster) Post(params *PostParams) error {
	ctx := context.Background()
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(p.endpoint),
		otlpmetricgrpc.WithHeaders(p.headers),
		otlpmetricgrpc.WithCompressor("gzip"),
	)
	if err != nil {
		return err
	}

	res := resource.NewWithAttributes(
		"",
		convertMapToAttrs(params.ResourceAttrs)...,
	)

	reader := sdkMetric.NewPeriodicReader(exporter)
	provider := sdkMetric.NewMeterProvider(
		sdkMetric.WithReader(reader),
		sdkMetric.WithResource(res),
	)
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
		log.Println("🚀 posted.")
	}()

	meter := provider.Meter(
		"github.com/Arthur1/otlc",
		/*
			metric.WithInstrumentationAttributes(
				convertMapToAttrs(params.ScopeAttrs)...,
			),
		*/
	)
	_, err = meter.Float64ObservableGauge(
		params.Name,
		metric.WithDescription(params.Description),
		metric.WithFloat64Callback(func(_ context.Context, o metric.Float64Observer) error {
			o.Observe(
				params.DataPointValue,
				metric.WithAttributes(
					convertMapToAttrs(params.DataPointAttrs)...,
				),
			)
			return nil
		}),
	)
	return err
}

func convertMapToAttrs(attrsMap map[string]string) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, len(attrsMap))
	for k, v := range attrsMap {
		attr := attribute.KeyValue{
			Key:   attribute.Key(k),
			Value: attribute.StringValue(v),
		}
		attrs = append(attrs, attr)
	}
	return attrs
}
