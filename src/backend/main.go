package main

import (
	_ "fly_easy/database"
	"fly_easy/routers"
	"net/http"
	_ "net/http"
)

func main() {
	// database.Tmp()
	http.ListenAndServe(":8080", routers.GetRouter())

}
