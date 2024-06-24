package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Arthur1/otlc/internal/metric"
	"github.com/Arthur1/otlc/internal/resource"
)

type MetricsPostCmd struct {
	OTLPEndpoint       string            `name:"otlp-endpoint" required:"" env:"OTEL_EXPORTER_OTLP_ENDPOINT,OTEL_EXPORTER_OTLP_METRICS_ENDPOINT" help:"OTLP endpoint"`
	OTLPHeaders        map[string]string `name:"otlp-headers" env:"OTEL_EXPORTER_OTLP_HEADERS,OTEL_EXPORTER_OTLP_METRICS_HEADERS" help:"OTLP headers"`
	OTLPProtocol       string            `name:"otlp-protocol" env:"OTEL_EXPORTER_OTLP_PROTOCOL" default:"grpc" enum:"grpc,http" help:"OTLP protocol"`
	OTLPInsecure       bool              `name:"otlp-insecure" help:"disable secure connection (required for such as localhost)"`
	MetricName         string            `name:"name" short:"n" required:"" help:"metric name"`
	MetricType         string            `name:"type" short:"t" default:"gauge" enum:"gauge,sum" help:"metric type"`
	MetricDescription  string            `name:"description" short:"d" help:"metric description"`
	MetricUnit         string            `name:"unit" short:"u" default:"1" help:"metric unit"`
	ResourceAttrs      map[string]string `name:"resource-attrs" mapsep:"," help:"resource attributes"`
	DataPointAttrs     map[string]string `name:"datapoint-attrs" mapsep:"," aliases:"attrs" help:"datapoint attributes"`
	DataPointTimestamp int64             `name:"timestamp" help:"datapoint timestamp (unix seconds)"`
	DataPointValue     float64           `arg:"" required:"" help:"datapoint value"`
}

func (c *MetricsPostCmd) Run(globals *Globals) error {
	var datapointTime time.Time
	if c.DataPointTimestamp == 0 {
		datapointTime = time.Now()
	} else {
		datapointTime = time.Unix(c.DataPointTimestamp, 0)
	}

	exporter, err := metric.NewExporter(&metric.NewExporterParams{
		OTLPEndpoint: c.OTLPEndpoint,
		OTLPHeaders:  c.OTLPHeaders,
		OTLPProtocol: c.OTLPProtocol,
		OTLPInsecure: c.OTLPInsecure,
	})
	if err != nil {
		return err
	}

	rsrc := resource.Generate(&resource.GenerateParams{
		ResourceAttrs: c.ResourceAttrs,
	})

	resourceMetrics, err := metric.Generate(&metric.GenerateParams{
		Resource:       rsrc,
		MetricName:     c.MetricName,
		MetricType:     c.MetricType,
		DataPointAttrs: c.DataPointAttrs,
		DataPointTime:  datapointTime,
		DataPointValue: c.DataPointValue,
	})
	if err != nil {
		return err
	}

	if err := exporter.Export(context.Background(), resourceMetrics); err != nil {
		return err
	}

	fmt.Println("exported.")

	return nil
}
