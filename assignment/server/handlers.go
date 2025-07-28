package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JeissonP09/todo"
)

// router configures routes to handle tasks.
// if the route is /todo, delegate to geatAllHandler or addHAndler
// If the route is /todo/:id and the method is GET, call getOneHandler to return a specific task.
func router(dataFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list todo.List

		err := list.Get(dataFile)
		if err != nil {
			errorReply(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		path := r.URL.Path
		//fmt.Println("path received:", path)

		if strings.HasPrefix(path, "/todo/") && r.Method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/todo/")
			//fmt.Println("idStr received:", idStr)

			id, err := validateID(idStr, &list)
			if err != nil {
				errorReply(w, r, http.StatusBadRequest, err.Error())
				return
			}
			getOneHandler(w, r, &list, id)
			return
		}

		if path == "/todo" {
			switch r.Method {
			case http.MethodGet:
				getAllHandler(w, r, &list)
			case http.MethodPost:
				addHandler(w, r, &list, dataFile)
			default:
				errorReply(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			}
			return
		}
		errorReply(w, r, http.StatusNotFound, "404 page not found")
	}
}

// getOneHandler responds with a specific task in JSON format,
// receives the task list and the previously validated ID,
// returns a JSON object with the requested task
func getOneHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int) {
	item := (*list)[id]
	reply := &todoResponse{
		Results: []todo.Todo{item},
	}
	jsonReply(w, r, http.StatusOK, reply)
}

// validateID converts from string to int, and validates that it is within the range of the task list.
func validateID(idStr string, list *todo.List) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(*list) {
		return 0, fmt.Errorf("invalid ID: out the range")
	}
	return id, nil
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
