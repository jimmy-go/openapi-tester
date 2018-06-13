package openapitester

import (
	"regexp"
)

var (
	rexVars = regexp.MustCompile(`\{.*?\}`)
)

func removeVars(uri string) string {
	return rexVars.ReplaceAllString(uri, "*")
}
