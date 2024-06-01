package handlers

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Gello")
}
