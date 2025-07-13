package handlers

import (
	"net/http"
)

func HelpHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello help!"))
}