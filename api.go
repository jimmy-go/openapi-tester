package openapitester

import (
	"errors"
	"fmt"
	"strings"
)

// API type contains all swagger info.
type API struct {
	Host        string                            `json:"host"`
	Paths       map[string]map[string]interface{} `json:"paths"`
	Definitions map[string]*Definition            `json:"definitions"`
}

// Search searchs method and request uri skipping url params as '/path/*/something'.
func (a *API) Search(method, requestURI string) (*PathMethod, error) {
	for k, uris := range a.Paths {
		kc := removeVars(k)
		rc := removeVars(requestURI)
		if kc == rc {
			for method2, v := range uris {
				p, ok := v.(*PathMethod)
				if !ok {
					continue
				}
				if strings.ToUpper(method2) == strings.ToUpper(method) {
					return p, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("not found: %s %s", method, requestURI)
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
