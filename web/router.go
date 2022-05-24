package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	ID        string `json:"id"`
}

func (u User) Validate() error {
	if u.FirstName == "" {
		return errors.New("Empty First Name")
	}
	if u.LastName == "" {
		return errors.New("Empty Last Name")
	}

	return nil
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	users := []User{User{FirstName: "john", LastName: "doe", ID: "1"}}

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		if err := user.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.ID = fmt.Sprintf("%d", len(users)+1)
		users = append(users, user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		params := mux.Vars(r)
		for _, user := range users {
			if user.ID == params["id"] {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}).Methods("GET")

	return router
}
