package handlers

import (
	_ "database/sql"
	"encoding/json"
	"fly_easy/database"
	"fly_easy/utils"
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
	id_str := r.URL.Query().Get("ID")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	email := r.URL.Query().Get("Email")
	if !(utils.CheckEmail(email)) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userByEmail := db.GetUserByEmail(email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userByEmail)
}

func GetSearchTickets(w http.ResponseWriter, r *http.Request) {
	depLocID, err1 := strconv.Atoi(r.URL.Query().Get("DeparteLocation"))
	arrLocID, err2 := strconv.Atoi(r.URL.Query().Get("ArriveLocation"))
	depDate := r.URL.Query().Get("DeparteDate")
	arrDate := r.URL.Query().Get("ArriveDate")
	isBusStr := r.URL.Query().Get("isBusiness")
	var isBusiness bool
	if err1 != nil && err2 != nil && !(utils.CheckDate(depDate) && utils.CheckDate(arrDate)) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isBusStr == "True" || isBusStr == "true" || isBusStr == "1" {
		isBusiness = true
	} else {
		isBusiness = false
	}
	serchedTockets := db.GetTicketsByCitesAndDate(depLocID, arrLocID, depDate, arrDate, isBusiness)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serchedTockets)
}
