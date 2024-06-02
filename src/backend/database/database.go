package database

import (
	"fmt"
	"log"
	"database/sql"
	"fly_easy/config"
  
	_ "github.com/go-sql-driver/mysql"
)


type User struct {
  ID            int
  Name          string
  LastName      string
  SurName       string
  Email         string
  PhoneNumber   int
};

type Ticket struct{ 
  ID            int
  DepartLocID   int
  ArriveLocID   int
  Price         int
  Airline       string
  DepTime       string
  DepDate       string
  ArriveTime    string
  ArriveDate    string
}

type Tickets []Ticket

type Location struct {
	ID         int
	Name       string
  Popularity float32
}

type LocationAndPrice struct {
	Name       string
  Price      int
}

type LocPrices []LocationAndPrice

type Locations []string

type IDataBase interface {
	Connect()
	Close()

  GetLocationsAndMinPrice() LocPrices
  GetPopularLocations() Locations

	GetUserByID(uid int) User
	GetUserByEmail(email string) User
  GetUserTicketsByID(uid int) Tickets
  GetUserFavoriteLocations(uid int) Locations

	GetTicketsByCitesAndDate() Tickets

  AddUser(user User, password string) bool
  AddTicketToFavorite(uid int) bool
  UpdateUserInfo(user User) bool
  UpdateUserPassword(uid int, password string) bool

  DeleteUser(id int) bool
  DeleteTicketFromFavorite() bool
}

type DB struct {
	Url string
	db  *sql.DB
}

func (d *DB) Connect() {
	var err error
	d.db, err = sql.Open("mysql", d.Url)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = d.db.Ping()
	if err != nil {
		log.Fatalf("Не удалось проверить подключение к базе данных: %v", err)
	}
}

func (d *DB) Close() {
	d.db.Close()
}


func (d *DB) GetLocationsAndMinPrice() LocPrices {
	db := d.db

  query := `
  SELECT l.LocName, MIN(f.price) AS min_price
  FROM Ticket f
  JOIN Location l ON f.ArriveLocID = l.ID
  GROUP BY l.LocName
  `

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data LocPrices

	for rows.Next() {
		var loc LocationAndPrice
		err := rows.Scan(&loc.Name, &loc.Price)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
    data = append(data, loc)
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data 
}

func (d *DB) GetPopularLocations() Locations {
	db := d.db

  query := `
  SELECT LocName FROM Location
  ORDER BY Popularity DESC;
  `

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data Locations

	for rows.Next() {
		var loc string
		err := rows.Scan(&loc)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
    data = append(data, loc)
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data
}
//
func (d *DB) GetUserByID(uid int) User {
	db := d.db

  query := `
  SELECT ID, Name, LastName, SurName, Email, PhoneNumber
  FROM User
  WHERE ID = ?
  `

	rows, err := db.Query(query, uid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data User

	for rows.Next() {
		err := rows.Scan(
      &data.ID, &data.Name, 
      &data.LastName, &data.SurName, 
      &data.Email, &data.PhoneNumber, 
      )
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data
}

func (d *DB) GetUserByEmail(email string) User {
	db := d.db

  query := `
  SELECT ID, Name, LastName, SurName, Email, PhoneNumber
  FROM User
  WHERE Email = ?
  `

	rows, err := db.Query(query, email)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data User

	for rows.Next() {
		err := rows.Scan(
      &data.ID, &data.Name, 
      &data.LastName, &data.SurName, 
      &data.Email, &data.PhoneNumber, 
      )
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data
}

func (d *DB) GetUserTicketsByID(uid int) Tickets {
	db := d.db

  query := `
  SELECT ID, DeparteLocID, ArriveLocID, Price, Airline,
  DepTime, DepDate
  FROM Ticket
  WHERE UserID = ?
  `

	rows, err := db.Query(query, uid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data Tickets

	for rows.Next() {
    var ticket Ticket
		err := rows.Scan(
      &ticket.ID,
      &ticket.DepartLocID, &ticket.ArriveLocID,
      &ticket.Price, &ticket.Airline,
      &ticket.DepTime, &ticket.DepDate,
      )
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
    data = append(data, ticket)
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data

}
func (d *DB) GetUserFavoriteLocations(uid int) Locations {
	db := d.db

  query := `
  SELECT l.LocName 
  FROM User u
  JOIN Favorites ul ON u.ID = ul.UserID
  JOIN Location l ON ul.LocationID = l.ID
  WHERE u.ID = ?
  `

	rows, err := db.Query(query, uid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  var data Locations

	for rows.Next() {
    var loc string
		err := rows.Scan(&loc)
		if err != nil {
			log.Fatalf("Не удалось считать данные: %v", err)
		}
    data = append(data, loc)
		// fmt.Printf("{%v, %v}\n",
  //     &loc.Name, &loc.Price,
  //   )
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

  return data
}

// func (d *DB) GetTicketsByCitesAndDate() Tickets {}

func (d *DB) AddUser(user User, passwordHash string) bool {
	db := d.db

  query := `
  INSERT INTO
  User(Name, Email, PasswordHash)
  VALUES
  (?, ?, ?)
  `

	result, err := db.Exec(query, user.Name, user.Email, passwordHash)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  if count, _ := result.RowsAffected(); count == 0 {
    return false
  }

  return true
}

func (d *DB) AddTicketToFavorite(uid, locid int) bool {
	db := d.db

  query := `
  INSERT INTO
  Favorites(UserID, LocationID)
  VALUES
  (?, ?)
  `

	result, err := db.Exec(query, uid, locid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  if count, _ := result.RowsAffected(); count == 0 {
    return false
  }

  return true
}

// func (d *DB) UpdateUserInfo(user User) bool {}
// func (d *DB) UpdateUserPassword(uid int, password string) bool {}
//
func (d *DB) DeleteUser(uid int) bool {
	db := d.db

  query := `
  DELETE FROM User
  WHERE ID = ?
  `

	result, err := db.Exec(query, uid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  if count, _ := result.RowsAffected(); count == 0 {
    return false
  }

  return true

}

func (d *DB) DeleteTicketFromFavorite(uid, locid int) bool {
	db := d.db

  query := `
  DELETE FROM Favorites
  WHERE UserID = ? AND LocationID = ?
  `

	result, err := db.Exec(query, uid, locid)
	if err != nil {
		log.Fatalf("Не удалось выполнить запрос: %v", err)
	}

  if count, _ := result.RowsAffected(); count == 0 {
    return false
  }

  return true
}


func Tmp() {
	db := DB{Url: config.DBUrl}
	db.Connect()

  // Get

  fmt.Println(db.GetLocationsAndMinPrice())
  fmt.Println(db.GetPopularLocations())
  fmt.Println(db.GetUserByID(4))
  fmt.Println(db.GetUserByEmail("emily.davis@example.com"))
  fmt.Println(db.GetUserTicketsByID(2))
  fmt.Println(db.GetUserFavoriteLocations(5))

  // Add
  // u := User{
  //   Name: "Bob",
  //   Email: "Bobasd@gmail.com",
  // }
  // fmt.Println(db.AddUser(u, "asdasdasdaklasd"))
  // fmt.Println(db.AddTicketToFavorite(11, 3))
  // fmt.Println(db.AddTicketToFavorite(11, 2))
  // fmt.Println(db.GetUserFavoriteLocations(11))
  // fmt.Println(db.DeleteTicketFromFavorite(11, 3))
  // fmt.Println(db.GetUserFavoriteLocations(11))
  fmt.Println(db.GetUserByID(11))
  fmt.Println(db.DeleteUser(11))
  fmt.Println(db.GetUserByID(11))



	db.Close() }
