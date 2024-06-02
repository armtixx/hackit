package handlers

import (
	_ "database/sql"
	"encoding/json"
	"fly_easy/database"
	"net/http"
	"strconv"
)

var db *database.DB = database.GetDB()

func GetLocationsList(w http.ResponseWriter, r *http.Request) {

	locationList := db.GetLocationsAndMinPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locationList)
}

func GetPopularLocations(w http.ResponseWriter, r *http.Request) {
	popLocList := db.GetPopularLocations()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(popLocList)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := 5
	userInfo := db.GetUserByID(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}

func GetUsersTickets(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("ID")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userTickets := db.GetUserTicketsByID(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userTickets)
}

func GetUsersFavorites(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("ID")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userFavorites := db.GetUserFavoriteLocations(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userFavorites)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := "qweasd@gmail.com"

	userByEmail := db.GetUserByEmail(email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userByEmail)
}

func GetSearchTickets(w http.ResponseWriter, r *http.Request) {
	// searchedTickets := db.GetTicketsByCitesAndDate()
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(searchedTickets)
}
