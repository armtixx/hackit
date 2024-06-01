package database

import (
	"database/sql"
	"fly_easy/config"
	"fmt"
	"log"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID          int
	Name        string
	LastName    string
	SurName     string
	Email       string
	PhoneNumber int
}

type Ticket struct {
	ID          int
	CustomerID  int
	DepartlocID int
	ArrivelocID int
	Price       int
	Airline     string
	Deptime     string
	DepDate     string
	Chcode      string
	TimeTaken   string
}

type Location struct {
	ID   int
	Name string
}

type IDataBase interface {
	Connect()
	Close()
	// Query()
	GetCustomer()
	GetTicket()
	GetLocation()
}

type DB struct {
	Url string
	db  *sql.DB
}

func (d *DB) Connect() {
	// Подключение к базе данных
	var err error
	d.db, err = sql.Open("mysql", d.Url)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Проверка подключения
	err = d.db.Ping()
	if err != nil {
		log.Fatalf("Не удалось проверить подключение к базе данных: %v", err)
	}
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) GetCustomer() {
	db := d.db
	query := "SELECT * FROM CUSTOMERS"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.ID, &customer.Name,
			&customer.LastName, &customer.SurName,
			&customer.Email, &customer.PhoneNumber,
		)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		fmt.Printf("{%v, %v, %v, %v, %v, %v}\n",
			customer.ID, customer.Name,
			customer.LastName, customer.SurName,
			customer.Email, customer.PhoneNumber,
		)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}
}

func (d *DB) GetTicket() {
	db := d.db
	query := "SELECT * FROM TICKET"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

	for rows.Next() {
		var ticket Ticket
		err := rows.Scan(
			&ticket.ID, &ticket.CustomerID,
			&ticket.DepartlocID, &ticket.ArrivelocID,
			&ticket.Price, &ticket.Airline,
			&ticket.Deptime, &ticket.DepDate,
			&ticket.Chcode, &ticket.TimeTaken,
		)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		fmt.Printf("{%v, %v,%v, %v, %v, %v, %v, %v, %v, %v}\n",
			ticket.ID, ticket.CustomerID,
			ticket.DepartlocID, ticket.ArrivelocID,
			ticket.Price, ticket.Airline,
			ticket.Deptime, ticket.DepDate,
			ticket.Chcode, ticket.TimeTaken,
		)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}
}

func (d *DB) GetLocation() {
	db := d.db
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

func Tmp() {
	db := DB{Url: config.DBUrl}
	db.Connect()
	db.GetTicket()
	// db.GetCustomer()
	// db.GetLocation()
	db.Close()
}
