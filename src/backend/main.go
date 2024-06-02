package main

import (
	"fly_easy/routers"
	_ "fly_easy/routers"
	_ "fmt"

	"net/http"
	_ "net/http"
)

func main() {
	r := routers.GetRouter()
	http.ListenAndServe(":8080", r)
}
