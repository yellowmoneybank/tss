package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Here are all trusted Shareholders declared, in a real world example, this has to be implemented differently!
var SHAREHOLDERS = []string{
	"localhost:8081",
	"localhost:8082",
	"localhost:8083",
}
var THRESHOLD = 2

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dealer")
	})
	// expects the secret as multipart/form-data POST, and returns the uuid of the secret
	api.HandleFunc("/secret", storeSecret).Methods(http.MethodPost)
	api.HandleFunc("/secret/{uuid}", getSecret).Methods(http.MethodGet)
	//	redistribute
	log.Fatalln(http.ListenAndServe(":8080", r))
}
