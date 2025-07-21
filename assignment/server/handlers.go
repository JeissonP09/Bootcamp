package main

import (
	"encoding/json"
	"net/http"

	"github.com/JeissonP09/todo"
)

func addHandler(w http.ResponseWriter, r *http.Request, list *todo.List, dataFile string) {
	type NewTask struct {
		Task string `json:"task"`
	}

	var item NewTask

	// Decode the JSON of the body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		errorReply(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Add new task to the list in the file JSON
	list.Add(item.Task)

	// Save the task
	err = list.Save(dataFile)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Response success with code 201
	textReply(w, r, http.StatusCreated, "Task created successfully")
}

func router(dataFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list todo.List

		// Loads the tasks from the file
		err := list.Get(dataFile)
		if err != nil {
			errorReply(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		// Decide which handler to use
		switch r.Method {
		case http.MethodGet:
			getAllHandler(w, r, &list)
		case http.MethodPost:
			addHandler(w, r, &list, dataFile)
		default:
			errorReply(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request, list *todo.List) {
	// Create the JSON response with the list of tasks in the field “results”
	reply := &todoResponse{Results: *list}
	// Send response in type json with status 200 Ok
	jsonReply(w, r, http.StatusOK, reply)
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
