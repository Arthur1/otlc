package metric

import (
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

type GenerateParams struct {
	Resource          *resource.Resource
	Scope             instrumentation.Scope
	MetricName        string
	MetricType        string
	MetricDescription string
	MetricUnit        string
	DataPointAttrs    map[string]string
	DataPointTime     time.Time
	DataPointValue    float64
}

func Generate(p *GenerateParams) (*metricdata.ResourceMetrics, error) {
	metrics := make([]metricdata.Metrics, 0, 1)
	attributes := make([]attribute.KeyValue, 0, len(p.DataPointAttrs))
	for k, v := range p.DataPointAttrs {
		attributes = append(attributes, attribute.String(k, v))
	}
	switch p.MetricType {
	case "gauge":
		metrics = append(metrics, metricdata.Metrics{
			Name:        p.MetricName,
			Description: p.MetricDescription,
			Unit:        p.MetricUnit,
			Data: metricdata.Gauge[float64]{
				DataPoints: []metricdata.DataPoint[float64]{
					{
						Time:       p.DataPointTime,
						Value:      p.DataPointValue,
						Attributes: attribute.NewSet(attributes...),
					},
				},
			},
		})
	case "sum":
		metrics = append(metrics, metricdata.Metrics{
			Name:        p.MetricName,
			Description: p.MetricDescription,
			Unit:        p.MetricUnit,
			Data: metricdata.Sum[float64]{
				IsMonotonic: true,
				Temporality: metricdata.CumulativeTemporality,
				DataPoints: []metricdata.DataPoint[float64]{
					{
						StartTime:  p.DataPointTime.Add(-1 * time.Minute),
						Time:       p.DataPointTime,
						Value:      p.DataPointValue,
						Attributes: attribute.NewSet(attributes...),
					},
				},
			},
		})
	default:
		return nil, fmt.Errorf("unexpected metric type: %s", p.MetricType)
	}

	rm := metricdata.ResourceMetrics{
		Resource: p.Resource,
		ScopeMetrics: []metricdata.ScopeMetrics{{
			Metrics: metrics,
			Scope:   p.Scope,
		}},
	}

	return &rm, nil

}
