package main

import (
  _ "fmt"

	_ "net/http"
	"fly_easy/database"
	_ "fly_easy/utils"
)

func main() {
	database.Tmp()
}
