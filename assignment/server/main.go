// Created main.go in directory Server
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// Created controller "rootHandler"
func rootHandler(w http.ResponseWriter, r *http.Request){
	// Validates if the path is correct
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World!!"))
}

func main() {
	// Created mux server
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	// Created flag "-p" type int, to manage the port
	port := flag.Int("p", 8080, "port of server")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(addr, mux))
}