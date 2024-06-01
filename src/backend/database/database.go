package database

import (
	"database/sql"
	"fly_easy/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Location struct {
	ID   int
	Name string
}

type iDataBase interface {
	Connect(DBName string)
	GetLocation()
	Close()
}

type DB struct {
	Url string
	db  *sql.DB
}

func (d *DB) Connect() {
	fmt.Println(d.Url)
	// Подключение к базе данных
	var err error
	d.db, err = sql.Open("mysql", d.Url)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	// defer d.db.Close()

	// Проверка подключения
	err = d.db.Ping()
	if err != nil {
		log.Fatalf("Не удалось проверить подключение к базе данных: %v", err)
	}

}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) GetLocation() {
	db := d.db
	rows, err := db.Query("SELECT ID, LOCATIONNAME FROM LOCATION")
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}
	//defer rows.Close()

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

func Tmp() {
	db := DB{Url: config.DBURL}
	db.Connect()
	db.GetLocation()
	db.Close()
}
