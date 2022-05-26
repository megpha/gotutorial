package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/megpha/website/data"
	"github.com/stretchr/testify/assert"
)

func TestAllUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestInvalidUserAdd(t *testing.T) {
	body, _ := json.Marshal(data.User{FirstName: "John", LastName: ""})
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}

func TestValidUserAdd(t *testing.T) {
	body, _ := json.Marshal(data.User{FirstName: "John", LastName: "Doe"})
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}

func TestFindUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var showUser data.User
	json.NewDecoder(response.Body).Decode(&showUser)

	assert.Equal(t, data.User{FirstName: "john", LastName: "doe", ID: "1"}, showUser)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}

func TestNotFoundRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/5", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}

func TestRemoveUnknownUser(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/users/5", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}

func TestRemoveKnownUser(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
}
