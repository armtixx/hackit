package handlers

import (
	"fmt"
	"net/http"
)

func GetLocationsList(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "loclist")

}
func GetPopularLocations(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "poploc")
}
func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "userinfo")
}
func GetUsersTickets(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "usertickets")
}
func GetUsersFavorites(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "favorites")
}
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "usbyemail")
}
func GetSearchTickets(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "searchtickets")
}
