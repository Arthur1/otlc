package scope

import (
	"testing"

	"github.com/Arthur1/otlc/internal/testutil"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in   *GenerateParams
		want instrumentation.Scope
	}{
		"generates empty instrumentation scope": {
			in:   &GenerateParams{},
			want: instrumentation.Scope{},
		},
		"generates an instrumentation scope with some settings": {
			in: &GenerateParams{
				ScopeName:      "github.com/Arthur1/otlc",
				ScopeVersion:   "0.0.0",
				ScopeSchemaURL: "https://opentelemetry.io/schemas/1.4.0",
			},
			want: instrumentation.Scope{
				Name:      "github.com/Arthur1/otlc",
				Version:   "0.0.0",
				SchemaURL: "https://opentelemetry.io/schemas/1.4.0",
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Generate(tt.in)
			testutil.NoDiff(t, tt.want, got, nil, "instrumentation scope should be equal but has diff")
		})
	}
}
