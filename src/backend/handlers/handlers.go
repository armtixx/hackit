package handlers

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func GetLocationsList(w http.ResponseWriter, r *http.Request)    {}
func GetPopularLocations(w http.ResponseWriter, r *http.Request) {}
func GetUserInfo(w http.ResponseWriter, r *http.Request)         {}
func GetUsersTickets(w http.ResponseWriter, r *http.Request)     {}
func GetUsersFavorites(w http.ResponseWriter, r *http.Request)   {}
func GetUserByEmail(w http.ResponseWriter, r *http.Request)      {}
func GetSearchTickets(w http.ResponseWriter, r *http.Request)    {}

// Add/
func AddUser(w http.ResponseWriter, r *http.Request)       {}
func AddToFavorite(w http.ResponseWriter, r *http.Request) {}

// Update/
func UpdateUserProfile(w http.ResponseWriter, r *http.Request)  {}
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {}

// Delete/
func DeleteUser(w http.ResponseWriter, r *http.Request)         {}
func DeleteFromFavorite(w http.ResponseWriter, r *http.Request) {}
