package main

import "net/http"

func newMux() http.Handler {
	// Created mux server
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	return mux
}