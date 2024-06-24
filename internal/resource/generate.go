package resource

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

type GenerateParams struct {
	ResourceAttrs map[string]string
}

func Generate(p *GenerateParams) *resource.Resource {
	attributes := make([]attribute.KeyValue, 0, len(p.ResourceAttrs))
	for k, v := range p.ResourceAttrs {
		attributes = append(attributes, attribute.String(k, v))
	}
	return resource.NewSchemaless(attributes...)
}
