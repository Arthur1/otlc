package scope

import (
	"testing"

	"github.com/Arthur1/otlc/internal/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
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
				ScopeAttrs:     map[string]string{"short_name": "otlc"},
			},
			want: instrumentation.Scope{
				Name:       "github.com/Arthur1/otlc",
				Version:    "0.0.0",
				SchemaURL:  "https://opentelemetry.io/schemas/1.4.0",
				Attributes: attribute.NewSet(attribute.String("short_name", "otlc")),
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Generate(tt.in)
			testutil.NoDiff(t, tt.want, got, []cmp.Option{cmpopts.IgnoreFields(instrumentation.Scope{}, "Attributes")}, "instrumentation scope should be equal but has diff")
			assert.True(t, tt.want.Attributes.Equals(&got.Attributes), "instrumentation scope attributes should be equal but has diff\n want: %+v\ngot: %+v", tt.want.Attributes.ToSlice(), got.Attributes.ToSlice())
		})
	}
}
