package openapitester

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	oapi := &API{}
	fnMap := map[string]func(string) string{
		"change_company": func(s string) string {
			return strings.Replace(s, "hi", "ho", -1)
		},
	}
	client := &http.Client{}
	ru, err := NewRunner(oapi, client, fnMap)
	assert.Nil(t, err)
	assert.NotNil(t, ru)
	report, err := ru.Exec(nil)
	assert.Nil(t, err)
	assert.NotNil(t, report)
}
