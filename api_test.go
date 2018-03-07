package openapitester

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	oapi := &API{
		Paths: map[string]map[string]*PathMethod{
			"/me/{userID}/profile": map[string]*PathMethod{
				"POST": &PathMethod{},
			},
		},
	}
	pm, err := oapi.Search("post", "/me/*/profile")
	assert.Nil(t, err)
	assert.NotNil(t, pm)

	pm2, err2 := oapi.Search("get", "/me/*/profile")
	assert.NotNil(t, err2)
	assert.Nil(t, pm2)
}

func TestExamples(t *testing.T) {
	oapi := &API{
		Paths: map[string]map[string]*PathMethod{
			"/login": map[string]*PathMethod{
				"POST": &PathMethod{
					Parameters: []Parameter{
						Parameter{
							Schema: &Schema{
								Ref: "#/definitions/Login",
							},
						},
					},
				},
			},
		},
		Definitions: map[string]*Definition{
			"Login": &Definition{
				Example: `{"username":"alice","password":"none"}`,
			},
		},
	}
	ss, err := oapi.Examples("post", "/login")
	assert.Nil(t, err)
	assert.NotNil(t, ss)

	exp := `{"username":"alice","password":"none"}`
	actual := ss[0]
	assert.EqualValues(t, exp, actual, actual)

	ss2, err2 := oapi.Examples("get", "/login")
	assert.NotNil(t, err2)
	assert.Nil(t, ss2)
}
