package handlers

import (
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello!!"))
}