package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// main sets the flags and launches the HTTP server.
func main() {
	dataFile := flag.String("f", "datafile.json", "data file of todo")
	host := flag.String("h", "localhost", "host fot default")
	port := flag.Int("p", 8080, "port of server")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Created the server with Address and Handler
	server := &http.Server{
		Addr:    addr,
		Handler: newMux(*dataFile),
	}

	log.Fatal(server.ListenAndServe())
}
