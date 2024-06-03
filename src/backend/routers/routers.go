package routers

import (
	"fly_easy/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Server", "FlyEasy")

		next.ServeHTTP(w, r)
	})
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/").Subrouter()
  api.Use(Middleware)

	apiGet := api.PathPrefix("/Get/").Subrouter()
	apiPost := api.PathPrefix("/Post/").Subrouter()
	apiDelete := api.PathPrefix("/Delete/").Subrouter()
	apiPut := api.PathPrefix("/Put/").Subrouter()

	apiGet.HandleFunc("/LocationsList", handlers.GetLocationsList)
	apiGet.HandleFunc("/PopularLocations", handlers.GetPopularLocations)
	apiGet.HandleFunc("/UserInfo", handlers.GetUserInfo)
	apiGet.HandleFunc("/UserTickets", handlers.GetUsersTickets)
	apiGet.HandleFunc("/UsersFavorites", handlers.GetUsersFavorites)
	apiGet.HandleFunc("/UserByEmail", handlers.GetUserByEmail)
	apiGet.HandleFunc("/SearchTickets", handlers.GetSearchTickets)

	apiPost.HandleFunc("/User", handlers.AddUser)
	apiPost.HandleFunc("/ToFavorite", handlers.AddToFavorite)

	apiPut.HandleFunc("/UserProfile", handlers.UpdateUserProfile)
	apiPut.HandleFunc("/UserPassword", handlers.UpdateUserPassword)

	apiDelete.HandleFunc("/User", handlers.DeleteUser)
	apiDelete.HandleFunc("/FromFavorite", handlers.DeleteFromFavorite)
	return r
}
