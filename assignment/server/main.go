package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// flags "-h" host and "-p" port
	host := flag.String("h", "localhost", "host fot default")
	port := flag.Int("p", 8080, "port of server")
	flag.Parse()

	// Construct the address (host:port)
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Created the server with Address and Handler
	server := &http.Server{
		Addr: addr,
		Handler: newMux(),
	}
	
	// Start the server and errors
	log.Fatal(server.ListenAndServe())
}