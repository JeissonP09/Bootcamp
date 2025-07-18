package main

import (
	"net/http"

	"github.com/JeissonP09/todo"
)

// Created a type with dataFile
type getAllHandler struct {
	dataFile string
}

func (h getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list := todo.List{}

	// We get data from the file
	err := list.Get(h.dataFile)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response
	resp := todoResponse{
		Results: list,
	}

	// Send response in type json with status 200 Ok
	jsonReply(w, r, http.StatusOK, &resp)

}

// Created controller "rootHandler"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Validates if the path is correct "/"
	// If there is an error we use the errorReply function
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "404 page not found")
		return
	}

	// Response with textReply function
	textReply(w, r, http.StatusOK, "Hello World!!")
}
