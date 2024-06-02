package handlers

import (
	"fmt"
	"net/http"
)

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "userprofile")
}
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "uspswd")
}
