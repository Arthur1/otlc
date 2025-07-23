package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
)

func TestNewExporter(t *testing.T) {
	t.Parallel()

	t.Run("returns otlpmetricgrpc.Exporter", func(t *testing.T) {
		t.Parallel()
		params := &NewExporterParams{
			OTLPProtocol: "grpc",
		}

		got, err := NewExporter(params)
		assert.NoError(t, err)
		assert.IsType(t, got, &otlpmetricgrpc.Exporter{})
		// TODO: test with dummy server
	})

	t.Run("returns otlpmetrichttp.Exporter", func(t *testing.T) {
		t.Parallel()
		params := &NewExporterParams{
			OTLPProtocol: "http",
		}

		got, err := NewExporter(params)
		assert.NoError(t, err)
		assert.IsType(t, got, &otlpmetrichttp.Exporter{})
		// TODO: test with dummy server
	})

	t.Run("returns otlpmetrichttp.Exporter (with scheme and custom path)", func(t *testing.T) {
		t.Parallel()
		params := &NewExporterParams{
			OTLPProtocol: "http",
			OTLPEndpoint: "http://localhost:4318/custom/path/v1/metrics",
		}

		got, err := NewExporter(params)
		assert.NoError(t, err)
		assert.IsType(t, got, &otlpmetrichttp.Exporter{})
	})

	t.Run("returns otlpmetrichttp.Exporter (no scheme, with custom path)", func(t *testing.T) {
		t.Parallel()
		params := &NewExporterParams{
			OTLPProtocol: "http",
			OTLPEndpoint: "localhost:4318/custom/path/v1/metrics",
		}

		got, err := NewExporter(params)
		assert.NoError(t, err)
		assert.IsType(t, got, &otlpmetrichttp.Exporter{})
	})

	t.Run("returns error when unexpected protocol is passed", func(t *testing.T) {
		t.Parallel()
		params := &NewExporterParams{
			OTLPProtocol: "unknown",
		}

		_, err := NewExporter(params)
		assert.Error(t, err)
	})
}
