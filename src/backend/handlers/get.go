package handlers

import (
	_ "database/sql"
	"encoding/json"
	"fly_easy/database"
	"net/http"
)

var db *database.DB = database.GetDB()

func GetLocationsList(w http.ResponseWriter, r *http.Request) {

	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
func GetPopularLocations(w http.ResponseWriter, r *http.Request) {
	PopLocList := db.GetPopularLocations()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PopLocList)
}
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
func GetUsersTickets(w http.ResponseWriter, r *http.Request) {
	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
func GetUsersFavorites(w http.ResponseWriter, r *http.Request) {
	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
func GetSearchTickets(w http.ResponseWriter, r *http.Request) {
	LocationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LocationList)
}
