package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"moritzm-mueller.de/tss/pkg/secretSharing"
	"moritzm-mueller.de/tss/pkg/shamir"
	"net/http"
)

func storeSecret(w http.ResponseWriter, r *http.Request) {
	// max value
	r.ParseMultipartForm(9223372036854775807)
	secret, _, err := r.FormFile("secret")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer secret.Close()

	secretBytes, err := ioutil.ReadAll(secret)
	if err != nil {
		fmt.Println(err)
		return
	}

	shares, err := shamir.SplitSecret(secretBytes, len(SHAREHOLDERS), uint8(THRESHOLD))
	if err != nil {
		fmt.Println(err)
		return
	}

	secretUUID := shares[0].ID
	ok := sendSharesTo(SHAREHOLDERS, shares)
	if ok {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"uuid":"`))
		_, _ = w.Write([]byte(secretUUID.String()))
		_, _ = w.Write([]byte(`"}`))
		return
	}

}
func getSecret(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	uuid, ok := queries["uuid"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
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
	b, err := json.Marshal(secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getSharesFromShareholder(shareholder string, uuid string) secretSharing.Share {
	//TODO
	return secretSharing.Share{}
}

func sendSharesTo(shareholders []string, shares []secretSharing.Share) bool {
	// TODO
	return false
}
