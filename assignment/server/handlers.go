package main

import "net/http"

// Created controller "rootHandler"
func rootHandler(w http.ResponseWriter, r *http.Request){
	// Validates if the path is correct "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Response
	w.Write([]byte("Hello World!!"))
}
