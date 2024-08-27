package resource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in   *GenerateParams
		want *resource.Resource
	}{
		"generates empty resource": {
			in: &GenerateParams{
				ResourceAttrs: map[string]string{},
			},
			want: resource.Empty(),
		},
		"generates resource with some attributes": {
			in: &GenerateParams{
				ResourceAttrs: map[string]string{"awesome.attribute": "hoge"},
			},
			want: resource.NewSchemaless(attribute.String("awesome.attribute", "hoge")),
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Generate(tt.in)
			assert.True(t, tt.want.Equal(got), "resource should be equal but has diff: %s", cmp.Diff(got, tt.want))
		})
	}
}
