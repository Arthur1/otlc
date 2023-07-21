package metrics

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkMetric "go.opentelemetry.io/otel/sdk/metric"

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

	reader := sdkMetric.NewPeriodicReader(exporter)
	provider := sdkMetric.NewMeterProvider(sdkMetric.WithReader(reader))
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
		log.Println("🚀 posted.")
	}()

	attrs := make([]attribute.KeyValue, 0, len(params.DataPointAttrs))
	for k, v := range params.DataPointAttrs {
		attr := attribute.KeyValue{
			Key:   attribute.Key(k),
			Value: attribute.StringValue(v),
		}
		attrs = append(attrs, attr)
	}

	meter := provider.Meter("github.com/Arthur1/otlc")
	_, err = meter.Float64ObservableGauge(
		params.Name,
		metric.WithDescription(params.Description),
		metric.WithFloat64Callback(func(_ context.Context, o metric.Float64Observer) error {
			attrs := metric.WithAttributes(attrs...)
			o.Observe(params.DataPointValue, attrs)
			return nil
		}),
	)
	return err
}
