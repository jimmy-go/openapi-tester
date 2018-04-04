package openapitester

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func defaultDo(c *http.Client, method, uri, payload string, headers map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBufferString(payload))
	if err != nil {
		return nil, 0, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
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
	DoFn       func(c *http.Client, method, uri, payload string, headers map[string]string) ([]byte, int, error)
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
	Method   string `json:"method"`
	Payload  []byte `json:"payload"`
	Response []byte `json:"response"`
	Error    string `json:"error"`
}

var (
	validMethods = map[string]bool{
		"GET":    true,
		"POST":   true,
		"PUT":    true,
		"PATCH":  true,
		"DELETE": true,
	}
)

// Exec runs a query against every endpoint registered.
func (r *Runner) Exec(headers map[string]string) ([]*Report, error) {
	var list []*Report
	for uri, pats := range r.API.Paths {
		// log.Printf("Exec : %s", uri)
		for method, pat := range pats {
			_ = pat
			if _, ok := validMethods[strings.ToUpper(method)]; !ok {
				// log.Printf("Exec : invalid method : %s uri : %s", method, uri)
				continue
			}

			var payload string
			// Get example payloads ONLY for GET methods.
			if strings.ToUpper(method) != "GET" {
				examples, err := r.API.Examples(method, uri)
				if err == nil {
					payload = examples[0]
				}
			}

			// Replace Request URI and payload vars.
			fullURL := r.API.Domain() + applyReplace(uri, r.ReplaceMap)
			payload = applyReplace(payload, r.ReplaceMap)

			// Do http request.
			var errMsg string
			res, code, err := r.DoFn(r.Client, method, fullURL, payload, headers)
			if err != nil {
				errMsg = err.Error()
			}
			// log.Printf("Exec : res : %s", res)

			// Replace body response.
			responseBody := applyReplace(string(res), r.ReplaceMap)

			re := &Report{
				Code:     code,
				Method:   method,
				Payload:  []byte(payload),
				Response: []byte(responseBody),
				Error:    errMsg,
			}
			list = append(list, re)
		}
	}
	return list, nil
}

func applyReplace(s string, fnMap map[string]func(string) string) string {
	if s == "" {
		return ""
	}
	for key, fn := range fnMap {
		ns := strings.Replace(s, key, fn(s), -1)
		s = ns
	}
	return s
}
