package handlers

import (
	"encoding/json"
	"fly_easy/database"
	"fmt"
	"net/http"
	// "strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	// id_str := r.URL.Query().Get("ID")
	// userID, _ := strconv.Atoi(id_str)
	user := database.User{
		Name:        r.URL.Query().Get("Name"),
		Email:       r.URL.Query().Get("Email"),
		PhoneNumber: r.URL.Query().Get("Phone"),
		SurName:     r.URL.Query().Get("SurName"),
		LastName:    r.URL.Query().Get("LastName"),
	}
	paswd := r.URL.Query().Get("password")
	serchedTockets, err := db.AddUser(user, paswd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serchedTockets)
}
func AddToFavorite(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "tofavor")
}
