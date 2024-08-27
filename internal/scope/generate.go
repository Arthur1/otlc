package scope

import "go.opentelemetry.io/otel/sdk/instrumentation"

type GenerateParams struct {
	ScopeName      string
	ScopeVersion   string
	ScopeSchemaURL string
}

func Generate(p *GenerateParams) instrumentation.Scope {
	s := instrumentation.Scope{
		Name:      p.ScopeName,
		Version:   p.ScopeVersion,
		SchemaURL: p.ScopeSchemaURL,
	}
	return s
}
