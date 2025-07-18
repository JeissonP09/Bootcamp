package main

import (
	"encoding/json"
	"time"

	"github.com/JeissonP09/todo"
)

type todoResponse struct {
	Results todo.List
}

func (tr todoResponse) MarshalJSON() ([]byte, error) {
	// Literal struct with struct tags
	response := struct {
		Results      todo.List `json:"results"`
		Date         time.Time `json:"date"`
		TotalResults int       `json:"total_results"`
	}{
		Results:      tr.Results,
		Date:         time.Now(),
		TotalResults: len(tr.Results),
	}
	return json.Marshal(response)
}
