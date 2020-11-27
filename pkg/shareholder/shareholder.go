package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const secretPath = "./secrets"

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dealer")
	})
	api.HandleFunc("/share", storeShare).Methods(http.MethodPost)
	api.HandleFunc("/share/{uuid}", getShare).Methods(http.MethodGet)
	api.HandleFunc("/share/{uuid}/redistribute", redistributeShare).Methods(http.MethodGet)
	api.HandleFunc("/redistshare/{uuid}", restoreRedistributeShare).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
