package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"moritzm-mueller.de/tss/pkg/secretSharing"
	"moritzm-mueller.de/tss/pkg/shamir"
)

// Here are all trusted Shareholders declared, in a real world example, this has to be implemented differently!
var SHAREHOLDERS = []string{
	"http://localhost:8081",
	// "localhost:8082",
	// "localhost:8083",
}
var THRESHOLD = 1

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dealer")
	})

	api.HandleFunc("/secret", storeSecret).Methods(http.MethodPost)
	api.HandleFunc("/secret/{uuid}", getSecret).Methods(http.MethodGet)
	//	redistribute
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func storeSecret(w http.ResponseWriter, r *http.Request) {
	secret := struct {
		Secret string `json:"secret"`
	}{}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	secretBytes := []byte(secret.Secret)

	shares, err := shamir.SplitSecret(secretBytes, len(SHAREHOLDERS), uint8(THRESHOLD))
	if err != nil {
		fmt.Println(err)

		return
	}

	secretUUID := shares[0].ID

	error := sendSharesTo(SHAREHOLDERS, shares)
	if error == nil {
		w.Header().Set("Content-Type", "application/json")

		response := struct {
			UUID string `json:"uuid"`
		}{
			UUID: secretUUID.String(),
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	println(error)
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	uuid, ok := queries["uuid"]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
		return
	}

	var shares []secretSharing.Share
	for _, shareholder := range SHAREHOLDERS {
		shares = append(shares, getSharesFromShareholder(shareholder, uuid))
	}

	secret, err := shamir.Reconstruct(shares)
	if err != nil {
		println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(string(secret))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getSharesFromShareholder(shareholder string, uuid string) secretSharing.Share {
	resp, err := http.Get(shareholder + "/api/share/" + uuid)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var share secretSharing.Share
	err = json.NewDecoder(resp.Body).Decode(&share)
	if err != nil {
		log.Fatalln(err)
	}

	return share
}

func sendSharesTo(shareholders []string, shares []secretSharing.Share) error {
	if len(shareholders) != len(shares) {
		return errors.New("Shareholders and shares length do not match")
	}

	for i, shareholder := range shareholders {
		requestBody, err := json.Marshal(shares[i])
		if err != nil {
			return errors.New("could not marshall share")
		}

		_, err = http.Post(shareholder+"/api/share", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalln("could not connect to shareholder")
		}
	}

	return nil
}
