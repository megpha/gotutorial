package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megpha/website/data"
)

var userStore = data.NewStorage()

func init() {
	userStore.Add(data.User{FirstName: "john", LastName: "doe", ID: "1"})
}

func jsonResponseType(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		handler(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userStore.All())
}

func create(w http.ResponseWriter, r *http.Request) {
	var user data.User
	json.NewDecoder(r.Body).Decode(&user)

	user, err := userStore.Add(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, error := userStore.Find(params["id"])

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func remove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, error := userStore.Remove(params["id"])

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", jsonResponseType(index)).Methods("GET")
	router.HandleFunc("/users", jsonResponseType(create)).Methods("POST")
	router.HandleFunc("/users/{id}", jsonResponseType(find)).Methods("GET")
	router.HandleFunc("/users/{id}", jsonResponseType(remove)).Methods("DELETE")
	return router
}
