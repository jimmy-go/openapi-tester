package openapitester

import (
	"log"
	"net/http"
)

func defaultDo(c *http.Client) ([]byte, error) {
	req, err := http.Request()
	return nil, nil
}

// Runner type.
type Runner struct {
	API    *API
	Client *http.Client
	DoFn   func(*http.Client) ([]byte, error)
}

// NewRunner returns a new runner that consumes oapi.
func NewRunner(oapi *API, client *http.Client, fnMap map[string]func(string) string) (*Runner, error) {
	ru := &Runner{
		API:    oapi,
		Client: client,
		DoFn:   defaultDo,
	}
	return ru, nil
}

// Exec runs a query against every endpoint registered.
func (r *Runner) Exec() error {
	for uri, pats := range r.API.Paths {
		log.Printf("Exec : %s", uri)
		for method, pat := range pats {
			payload, err := r.API.Examples(method, uri)
			if err != nil {
				return err
			}
			res, err := r.DoFn(r.client, method, uri, payload)
			if err != nil {
				return err
			}
			log.Printf("Exec : res : %s", res)
		}
	}
	return nil
}
