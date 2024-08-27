package metric

import (
	"testing"

	"github.com/Arthur1/otlc/internal/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in      *GenerateParams
		want    *metricdata.ResourceMetrics
		wantErr bool
	}{
		"generates gauge without any resource attribute and datapoint attribute": {
			in: &GenerateParams{
				Resource:          resource.Empty(),
				MetricName:        "awesome.gauge",
				MetricType:        "gauge",
				MetricDescription: "awesome gauge",
				MetricUnit:        "1",
				DataPointAttrs:    map[string]string{},
				DataPointTime:     testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue:    42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name:        "awesome.gauge",
						Description: "awesome gauge",
						Unit:        "1",
						Data: metricdata.Gauge[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
						},
					}},
				}},
			},
		},
		"generates gauge with some resource attributes": {
			in: &GenerateParams{
				Resource:       resource.NewSchemaless(attribute.String("service.name", "test service")),
				MetricName:     "awesome.gauge",
				MetricType:     "gauge",
				DataPointAttrs: map[string]string{},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.NewSchemaless(attribute.String("service.name", "test service")),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name: "awesome.gauge",
						Data: metricdata.Gauge[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
						},
					}},
				}},
			},
		},
		"generates gauge with an instrument scope": {
			in: &GenerateParams{
				Resource: resource.Empty(),
				Scope: instrumentation.Scope{
					Name: "github.com/Arthur1/otlc",
				},
				MetricName:     "awesome.gauge",
				MetricType:     "gauge",
				DataPointAttrs: map[string]string{},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Scope: instrumentation.Scope{
						Name: "github.com/Arthur1/otlc",
					},
					Metrics: []metricdata.Metrics{{
						Name: "awesome.gauge",
						Data: metricdata.Gauge[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
						},
					}},
				}},
			},
		},
		"generates gauge with some datapoint attributes": {
			in: &GenerateParams{
				Resource:       resource.Empty(),
				MetricName:     "awesome.gauge",
				MetricType:     "gauge",
				DataPointAttrs: map[string]string{"awesome.attribute": "hoge"},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name: "awesome.gauge",
						Data: metricdata.Gauge[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(attribute.String("awesome.attribute", "hoge")),
							}},
						},
					}},
				}},
			},
		},

		"generates sum without any resource attribute and datapoint attribute": {
			in: &GenerateParams{
				Resource:          resource.Empty(),
				MetricName:        "awesome.sum",
				MetricType:        "sum",
				MetricDescription: "awesome sum",
				MetricUnit:        "1",
				DataPointAttrs:    map[string]string{},
				DataPointTime:     testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue:    42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name:        "awesome.sum",
						Description: "awesome sum",
						Unit:        "1",
						Data: metricdata.Sum[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								StartTime:  testutil.Time(t, "2006-01-02T15:03:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
							IsMonotonic: true,
							Temporality: metricdata.CumulativeTemporality,
						},
					}},
				}},
			},
		},
		"generates sum with some resource attributes": {
			in: &GenerateParams{
				Resource:       resource.NewSchemaless(attribute.String("service.name", "test service")),
				MetricName:     "awesome.sum",
				MetricType:     "sum",
				DataPointAttrs: map[string]string{},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.NewSchemaless(attribute.String("service.name", "test service")),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name: "awesome.sum",
						Data: metricdata.Sum[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								StartTime:  testutil.Time(t, "2006-01-02T15:03:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
							IsMonotonic: true,
							Temporality: metricdata.CumulativeTemporality,
						},
					}},
				}},
			},
		},
		"generates sum with an instrument scope": {
			in: &GenerateParams{
				Resource: resource.Empty(),
				Scope: instrumentation.Scope{
					Name: "github.com/Arthur1/otlc",
				},
				MetricName:     "awesome.sum",
				MetricType:     "sum",
				DataPointAttrs: map[string]string{},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Scope: instrumentation.Scope{
						Name: "github.com/Arthur1/otlc",
					},
					Metrics: []metricdata.Metrics{{
						Name: "awesome.sum",
						Data: metricdata.Sum[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								StartTime:  testutil.Time(t, "2006-01-02T15:03:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(),
							}},
							IsMonotonic: true,
							Temporality: metricdata.CumulativeTemporality,
						},
					}},
				}},
			},
		},
		"generates sum with some datapoint attributes": {
			in: &GenerateParams{
				Resource:       resource.Empty(),
				MetricName:     "awesome.sum",
				MetricType:     "sum",
				DataPointAttrs: map[string]string{"awesome.attribute": "hoge"},
				DataPointTime:  testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue: 42,
			},
			want: &metricdata.ResourceMetrics{
				Resource: resource.Empty(),
				ScopeMetrics: []metricdata.ScopeMetrics{{
					Metrics: []metricdata.Metrics{{
						Name: "awesome.sum",
						Data: metricdata.Sum[float64]{
							DataPoints: []metricdata.DataPoint[float64]{{
								Time:       testutil.Time(t, "2006-01-02T15:04:05Z"),
								StartTime:  testutil.Time(t, "2006-01-02T15:03:05Z"),
								Value:      42,
								Attributes: attribute.NewSet(attribute.String("awesome.attribute", "hoge")),
							}},
							IsMonotonic: true,
							Temporality: metricdata.CumulativeTemporality,
						},
					}},
				}},
			},
		},

		"returns error when unexpected metric type is passed": {
			in: &GenerateParams{
				Resource:          resource.Empty(),
				MetricName:        "awesome.gauge",
				MetricType:        "histogram",
				MetricDescription: "awesome gauge",
				MetricUnit:        "1",
				DataPointAttrs:    map[string]string{},
				DataPointTime:     testutil.Time(t, "2006-01-02T15:04:05Z"),
				DataPointValue:    42,
			},
			wantErr: true,
		},
	}

	assertDataPointAttributes := func(t *testing.T, want, got *metricdata.ResourceMetrics) {
		switch wantData := want.ScopeMetrics[0].Metrics[0].Data.(type) {
		case metricdata.Gauge[float64]:
			gotData, ok := got.ScopeMetrics[0].Metrics[0].Data.(metricdata.Gauge[float64])
			assert.True(t, ok, "got is unexpected type: %T", got.ScopeMetrics[0].Metrics[0].Data)
			assert.Equal(t, wantData.DataPoints[0].Attributes, gotData.DataPoints[0].Attributes)
		case metricdata.Sum[float64]:
			gotData, ok := got.ScopeMetrics[0].Metrics[0].Data.(metricdata.Sum[float64])
			assert.True(t, ok, "got is unexpected type: %T", got.ScopeMetrics[0].Metrics[0].Data)
			assert.Equal(t, wantData.DataPoints[0].Attributes, gotData.DataPoints[0].Attributes)
		default:
			t.Errorf("want is unexpected type: %T", wantData)
		}
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Generate(tt.in)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				testutil.NoDiff(t, tt.want, got, []cmp.Option{cmpopts.IgnoreTypes(attribute.Set{})})
				assertDataPointAttributes(t, tt.want, got)
			}
		})
	}
}
