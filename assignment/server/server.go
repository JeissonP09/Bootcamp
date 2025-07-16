package main

import (
	"log"
	"net/http"
)

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