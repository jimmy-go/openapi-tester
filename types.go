package openapitester

import (
	"errors"
	"strings"
)

// API type contains all swagger info.
type API struct {
	Host        string                            `json:"host"`
	Paths       map[string]map[string]*PathMethod `json:"paths"`
	Definitions map[string]*Definition            `json:"definitions"`
}

// Search searchs method and request uri skipping url params as '/path/*/something'.
func (a *API) Search(method, requestURI string) (*PathMethod, error) {
	var pat map[string]*PathMethod
	for k, v := range a.Paths {
		kc := removeVars(k)
		rc := removeVars(requestURI)
		if kc == rc {
			pat = v
		}
	}
	m, ok := pat[strings.ToUpper(method)]
	if !ok {
		return nil, errors.New("method not found")
	}
	return m, nil
}

// Examples returns examples bodies if resource exists.
func (a *API) Examples(method, requestURI string) ([]string, error) {
	pm, err := a.Search(method, requestURI)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, x := range pm.Parameters {
		ref := strings.Replace(x.Schema.Ref, "#/definitions/", "", -1)
		def, ok := a.Definitions[ref]
		if !ok {
			continue
		}
		res = append(res, def.Example)
	}
	if len(res) == 0 {
		return nil, errors.New("example not found")
	}
	return res, nil
}

// PathMethod type.
type PathMethod struct {
	Tags        []string             `json:"tags"`
	Summary     string               `json:"summary"`
	Description string               `json:"description"`
	Consumes    []string             `json:"consumes,omitempty"`
	Parameters  []Parameter          `json:"parameters"`
	Responses   map[string]*Response `json:"responses"`
	Security    []interface{}        `json:"security,omitempty"` // TODO;
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
