package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// jsonReply responds with a JSON body and the given status code.
func jsonReply(w http.ResponseWriter, r *http.Request, status int, payload *todoResponse) {
	data, err := json.Marshal(payload)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(data)
}

// textReply responds with plain text and the given status code.
func textReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(status)

	w.Write([]byte(payload))
}

// errorReply prints the error and responds with the corresponding code and message.
func errorReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	log.Printf("Error %d: %s", status, payload)

	http.Error(w, payload, status)
}

// newMux creates and configures the HTTP router.
func newMux(dataFile string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.Handle("/todo", router(dataFile))

	return mux
}
