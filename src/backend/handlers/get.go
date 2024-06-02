package handlers

import (
	"database/sql"
	"encoding/json"
	"fly_easy/database"
	_ "fly_easy/database"
	"fmt"
	"net/http"
)

// var db database.IDataBase = database.DB{}

func GetLocationsList(w http.ResponseWriter, r *http.Request) {
	database.Tmp()

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
