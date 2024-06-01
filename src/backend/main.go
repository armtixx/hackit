package main

import (
	"fly_easy/database"
	_ "net/http"
)

func main() {
	database.Tmp()
}
