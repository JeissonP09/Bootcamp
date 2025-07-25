package main

import (
	"encoding/json"
	"net/http"

	"github.com/JeissonP09/todo"
)

// router configures routes to handle tasks.
func router(dataFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list todo.List

		err := list.Get(dataFile)
		if err != nil {
			errorReply(w, r, http.StatusInternalServerError, err.Error())
			return
		}

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

// addHAndler adds a new task to the file.
func addHandler(w http.ResponseWriter, r *http.Request, list *todo.List, dataFile string) {
	type NewTask struct {
		Task string `json:"task"`
	}

	var item NewTask

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		errorReply(w, r, http.StatusBadRequest, err.Error())
		return
	}

	list.Add(item.Task)

	err = list.Save(dataFile)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	textReply(w, r, http.StatusCreated, "Task created successfully")
}

// getAllHandler send all tasks as a JSON response.
func getAllHandler(w http.ResponseWriter, r *http.Request, list *todo.List) {
	reply := &todoResponse{Results: *list}
	jsonReply(w, r, http.StatusOK, reply)
}

// rootHandler response with a greeting if the URL is "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "404 page not found")
		return
	}

	textReply(w, r, http.StatusOK, "Hello World!!")
}
