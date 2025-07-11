package handlers

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello!!")
}