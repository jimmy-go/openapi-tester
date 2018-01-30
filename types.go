package openapitester

// API type contains all swagger info.
type API struct {
	Host        string                 `json:"host"`
	Paths       map[string]*Path       `json:"paths"`
	Definitions map[string]*Definition `json:"definitions"`
}

// Path type.
type Path struct {
	Get        *PathMethod   `json:"get,omitempty"`
	Post       *PathMethod   `json:"post,omitempty"`
	Put        *PathMethod   `json:"put,omitempty"`
	Delete     *PathMethod   `json:"delete,omitempty"`
	Parameters []*Parameters `json:"parameters,omitempty"`
}

// PathMethod type.
type PathMethod struct {
	Tags        []string             `json:"tags"`
	Summary     string               `json:"summary"`
	Description string               `json:"description"`
	Consumes    []string             `json:"consumes,omitempty"`
	Parameters  []Parameters         `json:"parameters"`
	Responses   map[string]*Response `json:"responses"`
	Security    []interface{}        `json:"security,omitempty"` // TODO;
}

// Parameters type.
type Parameters struct {
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
