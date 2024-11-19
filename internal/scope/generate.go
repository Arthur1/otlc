package scope

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

type GenerateParams struct {
	ScopeName      string
	ScopeVersion   string
	ScopeSchemaURL string
	ScopeAttrs     map[string]string
}

func Generate(p *GenerateParams) instrumentation.Scope {
	attributes := make([]attribute.KeyValue, 0, len(p.ScopeAttrs))
	for k, v := range p.ScopeAttrs {
		attributes = append(attributes, attribute.String(k, v))
	}
	s := instrumentation.Scope{
		Name:       p.ScopeName,
		Version:    p.ScopeVersion,
		SchemaURL:  p.ScopeSchemaURL,
		Attributes: attribute.NewSet(attributes...),
	}
	return s
}
