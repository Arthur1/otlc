package metric

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type NewExporterParams struct {
	OTLPEndpoint string
	OTLPHeaders  map[string]string
	OTLPProtocol string
	OTLPInsecure bool
}

func NewExporter(p *NewExporterParams) (sdkmetric.Exporter, error) {
	switch p.OTLPProtocol {
	case "grpc":
		options := []otlpmetricgrpc.Option{
			otlpmetricgrpc.WithEndpoint(p.OTLPEndpoint),
		}
		if len(p.OTLPHeaders) > 0 {
			options = append(options, otlpmetricgrpc.WithHeaders(p.OTLPHeaders))
		}
		if p.OTLPInsecure {
			options = append(options, otlpmetricgrpc.WithInsecure())
		}
		exporter, err := otlpmetricgrpc.New(context.Background(), options...)
		return exporter, err
	case "http":
		options := []otlpmetrichttp.Option{
			otlpmetrichttp.WithEndpoint(p.OTLPEndpoint),
		}
		if len(p.OTLPHeaders) > 0 {
			options = append(options, otlpmetrichttp.WithHeaders(p.OTLPHeaders))
		}
		if p.OTLPInsecure {
			options = append(options, otlpmetrichttp.WithInsecure())
		}
		exporter, err := otlpmetrichttp.New(context.Background(), options...)
		return exporter, err
	default:
		return nil, fmt.Errorf("unexpected protocol: %s", p.OTLPProtocol)
	}
}
