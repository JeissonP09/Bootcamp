package main

import "net/http"

// Created controller "rootHandler"
func rootHandler(w http.ResponseWriter, r *http.Request){
	// Validates if the path is correct "/"
	// If there is an error we use the errorReply function
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "404 page not found")
		return
	}

	// Response with textReply function
	textReply(w, r, http.StatusOK, "Hello World!!")
}
