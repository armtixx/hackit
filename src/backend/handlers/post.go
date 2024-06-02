package handlers

import (
	"fmt"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "useradd")
}
func AddToFavorite(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "tofavor")
}
