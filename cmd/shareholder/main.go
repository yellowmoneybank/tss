package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"moritzm-mueller.de/tss/pkg/secretSharing"
)

const secretPath = "/secrets"

var uniquePath string

func main() {
	// Each shareholder stores his shares in a unique folder
	suffix := uuid.New().String()
	path, _ := os.Getwd()
	uniquePath = path + secretPath + suffix

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "shareholder")
	})
	api.HandleFunc("/share", storeShare).Methods(http.MethodPost)
	api.HandleFunc("/share/{uuid}", getShare).Methods(http.MethodGet)

	// These are hard to implement in an meaningful way without knowing the
	// concrete infrastructure the wallet will run on. The get some insight
	// into how the redistribution process works, refer to the benchmarks
	// and the unit tests.

	// api.HandleFunc("/share/{uuid}/redistribute", redistributeShare).Methods(http.MethodGet)
	// api.HandleFunc("/redistshare/{uuid}", restoreRedistributeShare).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8081", r))
}

func getShare(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	id, ok := queries["uuid"]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
		return
	}

	var share secretSharing.Share

	// look for file in folder
	err := filepath.Walk(uniquePath, func(path string, info os.FileInfo, err error) error {
		if info.Name() == id+".share" {
			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			defer f.Close()
			err = json.NewDecoder(f).Decode(&share)
			if err != nil {
				http.Error(w, err.Error(), 400)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(share)
}

func storeShare(w http.ResponseWriter, r *http.Request) {
	var share secretSharing.Share

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&share)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	// create dir, if not exists

	os.MkdirAll(uniquePath, os.ModePerm)

	f, err := os.Create(uniquePath + "/" + share.ID.String() + ".share")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer f.Close()
	// it would be better to do this in binary
	s, err := json.Marshal(share)
	if err != nil {
		fmt.Println(err)

		return
	}

	_, err = f.Write(s)

	w.WriteHeader(http.StatusOK)
}
