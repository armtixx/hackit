package handlers

import (
	"fmt"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "deluser")
}
func DeleteFromFavorite(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "fromfavor")
}
