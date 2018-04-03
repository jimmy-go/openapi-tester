package openapitester

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func defaultDo(c *http.Client, method, uri, payload string) ([]byte, int, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBufferString(payload))
	if err != nil {
		return nil, 0, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("invalid response: %d", resp.StatusCode)
	}
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, 0, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, 0, err
	}
	return buf.Bytes(), 0, nil
}

// Runner type.
type Runner struct {
	API        *API
	Client     *http.Client
	DoFn       func(c *http.Client, method, uri, payload string) ([]byte, int, error)
	ReplaceMap map[string]func(string) string
}

// NewRunner returns a new runner that consumes oapi.
func NewRunner(oapi *API, client *http.Client, fnMap map[string]func(string) string) (*Runner, error) {
	ru := &Runner{
		API:        oapi,
		Client:     client,
		DoFn:       defaultDo,
		ReplaceMap: fnMap,
	}
	return ru, nil
}

// Report type.
type Report struct {
	Code     int    `json:"code"`
	Payload  []byte `json:"payload"`
	Response []byte `json:"response"`
}

// Exec runs a query against every endpoint registered.
func (r *Runner) Exec() ([]*Report, error) {
	var list []*Report
	for uri, pats := range r.API.Paths {
		log.Printf("Exec : %s", uri)
		for method, pat := range pats {
			_ = pat
			// Get example payloads.
			examples, err := r.API.Examples(method, uri)
			if err != nil {
				return nil, err
			}

			// Replace Request URI and payload vars.
			payload := applyReplace(examples[0], r.ReplaceMap)
			uri = applyReplace(uri, r.ReplaceMap)

			// Do http request.
			res, code, err := r.DoFn(r.Client, method, uri, payload)
			if err != nil {
				return nil, err
			}
			log.Printf("Exec : res : %s", res)

			// Replace body response.
			responseBody := applyReplace(string(res), r.ReplaceMap)

			re := &Report{
				Code:     code,
				Payload:  []byte(payload),
				Response: []byte(responseBody),
			}
			list = append(list, re)
		}
	}
	return list, nil
}

func applyReplace(s string, fnMap map[string]func(string) string) string {
	for key, fn := range fnMap {
		ns := strings.Replace(s, key, fn(s), -1)
		s = ns
	}
	return s
}
