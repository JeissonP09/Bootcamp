package handlers

import (
	"fmt"
	"net/http"
)

func Help(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello help!")
}