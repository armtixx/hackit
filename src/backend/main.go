package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Location struct {
	ID   int
	Name string
}

func main() {
	// Настройки подключения к базе данных
	dsn := "root:Aa12344321@tcp(5.39.249.80:3307)/FLYEASY"

	// Подключение к базе данных
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось проверить подключение к базе данных: %v", err)
	}

	// Выполнение запроса
	rows, err := db.Query("SELECT ID, LOCATIONNAME FROM LOCATION")
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}
	defer rows.Close()

	// Обработка результата
	for rows.Next() {
		var location Location
		err := rows.Scan(&location.ID, &location.Name)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", location.ID, location.Name)
	}

	// Проверка на ошибки после завершения чтения
	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}
}
