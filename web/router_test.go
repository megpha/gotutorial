package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	if http.StatusOK != response.Code {
		t.Errorf("Invalid response code")
	}
}

func TestAddUser(t *testing.T) {
	user := User{FirstName: "adding", LastName: "user"}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	if http.StatusCreated != response.Code {
		t.Errorf("Invalid response code")
	}
	var createdUser User
	json.NewDecoder(response.Body).Decode(&createdUser)
	if createdUser.ID == "" {
		t.Errorf("ID is missing")
	}
}

func TestUserRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	if http.StatusOK != response.Code {
		t.Errorf("Invalid response code")
	}
}

func TestUserNotFoundRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/2", nil)
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, req)

	if http.StatusNotFound != response.Code {
		t.Errorf("Invalid response code")
	}
}
