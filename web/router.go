package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var userStore = UserStore{Users: []User{{FirstName: "john", LastName: "doe", ID: "1"}}}

type UserStore struct {
	Users []User
}

func (u UserStore) Add(user User) (UserStore, error) {

	if error := user.Validate(); error != nil {
		return UserStore{Users: u.Users}, errors.New("Invalid User Information")
	}

	user.ID = fmt.Sprintf("%d", len(u.Users)+1)
	users := append(u.Users, user)

	return UserStore{Users: users}, nil
}

func (u UserStore) Find(id string) (User, error) {
	for _, user := range u.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.New("No User Found")
}

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
func LastUser(u UserStore) (User, error) {
	if len(u.Users) == 0 {
		return User{}, errors.New("empty store")
	}

	return u.Users[len(u.Users)-1], nil
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(userStore.Users)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	userStore, err := userStore.Add(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = LastUser(userStore)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := mux.Vars(r)
	user, error := userStore.Find(params["id"])

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", index).Methods("GET")
	router.HandleFunc("/users", create).Methods("POST")
	router.HandleFunc("/users/{id}", find).Methods("GET")
	return router
}
