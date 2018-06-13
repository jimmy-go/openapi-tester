package openapitester

// PathMethod type.
type PathMethod struct {
	Tags        []string                 `json:"tags"`
	Summary     string                   `json:"summary"`
	Description string                   `json:"description"`
	Consumes    []string                 `json:"consumes,omitempty"`
	Parameters  []*Parameter             `json:"parameters,omitempty"`
	Responses   map[string]*Response     `json:"responses,omitempty"`
	Security    []map[string]interface{} `json:"security,omitempty"` // TODO;
}

// Parameter  type.
type Parameter struct {
	Name     string  `json:"name"`
	In       string  `json:"in"`
	Required bool    `json:"required"`
	Type     string  `json:"type"`
	Schema   *Schema `json:"schema,omitempty"`
}

// Schema type.
type Schema struct {
	Ref string `json:"$ref"`
}

// Response type.
type Response struct {
	Description string  `json:"description"`
	Schema      *Schema `json:"schema,omitempty"`
}

// Property type.
type Property struct {
	Type string `json:"type"`
}

// Definition type.
type Definition struct {
	Type       string               `json:"type"`
	Properties map[string]*Property `json:"properties"`
	Example    string               `json:"example"`
}
