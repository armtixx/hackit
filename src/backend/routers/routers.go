package routers

import (
	"fly_easy/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set global headers here
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Server", "FlyEasy")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func GetRouter() *mux.Router {

	r := mux.NewRouter()

	api := r.PathPrefix("/api/").Subrouter()
	apiGet := api.PathPrefix("/Get/").Subrouter()
	apiPost := api.PathPrefix("/Post/").Subrouter()
	apiDelete := api.PathPrefix("/Delete/").Subrouter()
	apiPut := api.PathPrefix("/Put/").Subrouter()

	api.HandleFunc("/GetLocation", Middleware(handlers.Handler))

	apiGet.HandleFunc("/LocationsList", Middleware(handlers.GetLocationsList))
	apiGet.HandleFunc("/PopularLocations", Middleware(handlers.GetPopularLocations))
	apiGet.HandleFunc("/UserInfo", Middleware(handlers.GetUserInfo))
	apiGet.HandleFunc("/UserTickets", Middleware(handlers.GetUsersTickets))
	apiGet.HandleFunc("/UsersFavorites", Middleware(handlers.GetUsersFavorites))
	apiGet.HandleFunc("/UserByEmail", Middleware(handlers.GetUserByEmail))
	apiGet.HandleFunc("/SearchTickets", Middleware(handlers.GetSearchTickets))

	apiPost.HandleFunc("/User", Middleware(handlers.AddUser))
	apiPost.HandleFunc("/ToFavorite", Middleware(handlers.AddToFavorite))

	apiPut.HandleFunc("/UserProfile", Middleware(handlers.UpdateUserProfile))
	apiPut.HandleFunc("/UserPassword", Middleware(handlers.UpdateUserPassword))

	apiDelete.HandleFunc("/User", Middleware(handlers.DeleteUser))
	apiDelete.HandleFunc("/FromFavorite", Middleware(handlers.DeleteFromFavorite))
	return r
}
