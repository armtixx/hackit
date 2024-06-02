package database

import (
	"database/sql"
	"fly_easy/config"
	"fmt"
	"log"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

func GetLocationsList() {
	db := DB{Url: config.DBUrl}
	db.Connect()
	query := "SELECT ID, LOCATIONNAME FROM LOCATION"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

	for rows.Next() {
		var location Location
		err := rows.Scan(&location.ID, &location.Name)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", location.ID, location.Name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}
}
