package main

import (
	"fly_easy/database"
	"fly_easy/routers"
	"net/http"
	_ "net/http"
)

func main() {
	database.DataBaseInit()
	http.ListenAndServe(":8080", routers.GetRouter())

}
