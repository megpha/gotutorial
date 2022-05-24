package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestAddUser(t *testing.T) {
	type TestCase struct {
		TestUser User
		Code     int
	}
	testcases := []TestCase{
		TestCase{TestUser: User{FirstName: "John", LastName: ""}, Code: http.StatusBadRequest},
		TestCase{TestUser: User{FirstName: "", LastName: "Doe"}, Code: http.StatusBadRequest},
		TestCase{TestUser: User{FirstName: "", LastName: ""}, Code: http.StatusBadRequest},
		TestCase{TestUser: User{FirstName: "John", LastName: "Doe"}, Code: http.StatusCreated},
	}
	for _, ts := range testcases {
		body, _ := json.Marshal(ts.TestUser)
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
		response := httptest.NewRecorder()
		CreateRouter().ServeHTTP(response, req)
		assert.Equal(t, ts.Code, response.Code)
		if ts.Code == http.StatusCreated {
			var createdUser User
			json.NewDecoder(response.Body).Decode(&createdUser)
			assert.NotEqual(t, createdUser.ID, "")
		}
	}
}

func TestValidUser(t *testing.T) {
	testcases := []User{
		User{FirstName: "John", LastName: ""},
		User{FirstName: "", LastName: "Doe"},
		User{FirstName: "", LastName: ""},
	}

	for _, ts := range testcases {
		assert.NotNil(t, ts.Validate())
	}
}

func TestUserRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var showUser User
	json.NewDecoder(response.Body).Decode(&showUser)

	assert.Equal(t, User{FirstName: "john", LastName: "doe", ID: "1"}, showUser)
}

func TestUserNotFoundRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/2", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)
	assert.Equal(t, http.StatusNotFound, response.Code)
}
