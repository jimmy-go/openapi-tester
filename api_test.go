package openapitester

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	oapi := &API{
		Paths: map[string]map[string]json.RawMessage{
			"/me/{userID}/profile": map[string]json.RawMessage{
				"POST": json.RawMessage(` { "tags": [ "Schedules" ], "summary": "Add schedule array to route.", "consumes": [], "parameters": [ { "name": "Authorization", "in": "header", "required": false, "type": "string" }, { "name": "body", "in": "body", "required": true, "schema": { "$ref": "#/definitions/RequestSchedule" } } ], "responses": { "200": { "description": "Status 200", "schema": { "$ref": "#/definitions/ScheduleCreated" } } }, "security": [ { "JWT": [] } ] } `),
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
		Paths: map[string]map[string]json.RawMessage{
			"/login": map[string]json.RawMessage{
				"POST": json.RawMessage(` { "tags": [ "Schedules" ], "summary": "Add schedule array to route.", "consumes": [], "parameters": [ { "name": "Authorization", "in": "header", "required": false, "type": "string" }, { "name": "body", "in": "body", "required": true, "schema": { "$ref": "#/definitions/Login" } } ], "responses": { "200": { "description": "Status 200", "schema": { "$ref": "#/definitions/ScheduleCreated" } } }, "security": [ { "JWT": [] } ] } `),
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
	if err != nil {
		return
	}

	exp := `{"username":"alice","password":"none"}`
	actual := ss[0]
	assert.EqualValues(t, exp, actual, actual)

	ss2, err2 := oapi.Examples("get", "/login")
	assert.NotNil(t, err2)
	assert.Nil(t, ss2)
}
