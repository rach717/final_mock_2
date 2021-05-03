package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

var a App

func TestMain(m *testing.M) {
	a.Initialise()
	code := m.Run()
	os.Exit(code)
}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}
func TestGetNonexistentproduct(t *testing.T) {
	assert := assert.New(t)
	req, _ := http.NewRequest("GET", "/blueprints/1", nil)
	response := ExecuteRequest(req)
	assert.Equal(http.StatusBadRequest, response.Code)
}
func TestGetBlueprint(t *testing.T) {
	assert := assert.New(t)
	requ, _ := http.NewRequest("GET", "/blueprints/1234", nil)
	respo := ExecuteRequest(requ)
	var actual Blueprint
	_ = json.NewDecoder(respo.Body).Decode(&actual)
	//fetching from db
	expected := Blueprints[0]

	t.Logf("expected%T.actual%T", expected, actual)
	t.Logf("hello there..%v", reflect.DeepEqual(actual, expected))
	t.Log(actual)
	t.Log(expected)
	if reflect.DeepEqual(actual, expected) != true {
		t.Errorf("json fetched not equal to actual one")
	}
	assert.Equal(http.StatusOK, respo.Code)
}
func TestGetBlueprints(t *testing.T) {
	assert := assert.New(t)
	requ, _ := http.NewRequest("GET", "/blueprints", nil)
	respo := ExecuteRequest(requ)
	assert.Equal(http.StatusOK, respo.Code)
}
func TestAddBlueprint(t *testing.T) {
	assert := assert.New(t)
	var jsonStr = []byte(`{}`)
	req, _ := http.NewRequest("POST", "/blueprint", bytes.NewBuffer(jsonStr))
	response := ExecuteRequest(req)
	assert.Equal(http.StatusOK, response.Code)
}
