package routers

import (
	"fly_easy/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {

	r := mux.NewRouter()

	api := r.PathPrefix("/api/").Subrouter()
	apiGet := api.PathPrefix("/Get/").Subrouter()
	apiPost := api.PathPrefix("/Post/").Subrouter()
	apiDelete := api.PathPrefix("/Delete/").Subrouter()
	apiPut := api.PathPrefix("/Put/").Subrouter()
	api.HandleFunc("/GetLocation", handlers.Handler)

	apiGet.HandleFunc("LocationsList", handlers.GetLocationsList)
	apiGet.HandleFunc("UserInfo", handlers.GetUserInfo)
	apiGet.HandleFunc("UserTickets", handlers.GetUsersTickets)
	apiGet.HandleFunc("UsersFavorites)", handlers.GetUsersFavorites)
	apiGet.HandleFunc("UserByEmail)", handlers.GetUserByEmail)
	apiGet.HandleFunc("SearchTickets)", handlers.GetSearchTickets)

	apiPost.HandleFunc("User)", handlers.AddUser)
	apiPost.HandleFunc("ToFavorite)", handlers.AddToFavorite)

	apiPut.HandleFunc("UserProfile)", handlers.UpdateUserProfile)
	apiPut.HandleFunc("UserPassword)", handlers.UpdateUserPassword)

	apiDelete.HandleFunc("User)", handlers.DeleteUser)
	apiDelete.HandleFunc("FromFavorite)", handlers.DeleteFromFavorite)
	return r
}
