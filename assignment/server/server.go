package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonReply(w http.ResponseWriter, r *http.Request, status int, payload *todoResponse) {
	// Convert payload a format JSON and use func errorReply if fail
	data, err := json.Marshal(payload)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// sets the type of content and format JSON
	w.Header().Set("Content-Type", "application/json")

	// sets the HTTP status code
	w.WriteHeader(status)

	// send the JSON
	w.Write(data)
}

func textReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	// sets the type of content and text plain
	w.Header().Set("Content-Type", "text/plain")

	// sets the HTTP status code
	w.WriteHeader(status)

	// write the menssage in the body of the response
	w.Write([]byte(payload))
}

func errorReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	// print error in console
	log.Printf("Error %d: %s", status, payload)

	// send response of error to the client
	http.Error(w, payload, status)
}

func newMux() http.Handler {
	// Created mux server
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	return mux
}