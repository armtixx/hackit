package database

import (
	"database/sql"
	"fly_easy/config"
	"fly_easy/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int
	Name        string
	LastName    string
	SurName     string
	Email       string
	PhoneNumber string
}

type Ticket struct {
	ID          int
	DepartLocID int
	ArriveLocID int
	Price       int
	Airline     string
	DepTime     string
	DepDate     string
	ArriveTime  string
	ArriveDate  string
}

type Location struct {
	ID         int
	Name       string
	Popularity float32
}

type LocationAndPrice struct {
	Name  string
	Price int
}

type Tickets []Ticket
type LocPrices []LocationAndPrice
type Locations []string

type DB struct {
	url string
	db  *sql.DB
}

func (d *DB) Connect() error {
	var err error = nil
	d.db, err = sql.Open("mysql", d.url)
	if err != nil {
		return err
	}
	if err = d.db.Ping(); err != nil {
		return err
	}
	return err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetLocationsAndMinPrice() (LocPrices, error) {
	db := d.db
	query :=
		`
  SELECT l.LocName, MIN(f.price) AS min_price
  FROM Ticket f
  JOIN Location l ON f.ArriveLocID = l.ID
  GROUP BY l.LocName
  `

	rows, err := db.Query(query)
	if err != nil {
		return LocPrices{}, err
	}

	var data LocPrices
	for rows.Next() {
		var loc LocationAndPrice
		err := rows.Scan(&loc.Name, &loc.Price)
		if err != nil {
			return LocPrices{}, err
		}

		data = append(data, loc)
	}

	err = rows.Err()
	if err != nil {
		return LocPrices{}, err
	}

	return data, nil
}

func (d *DB) GetPopularLocations() (Locations, error) {
	db := d.db
	query :=
		`
  SELECT LocName FROM Location
  ORDER BY Popularity DESC;
  `

	rows, err := db.Query(query)
	if err != nil {
		return Locations{}, err
	}

	var data Locations
	for rows.Next() {
		var loc string
		err := rows.Scan(&loc)
		if err != nil {
			return Locations{}, err
		}
		data = append(data, loc)
	}

	err = rows.Err()
	if err != nil {
		return Locations{}, err
	}

	return data, nil
}

func (d *DB) GetUserByID(uid int) (User, error) {
	db := d.db
	query :=
		`
  SELECT ID, Name, LastName, SurName, Email, PhoneNumber
  FROM User
  WHERE ID = ?
  `

	row := db.QueryRow(query, uid)
	var data User
	err := row.Scan(
		&data.ID, &data.Name,
		&data.LastName, &data.SurName,
		&data.Email, &data.PhoneNumber,
	)
	if err != nil {
		return User{}, err
	}

	err = row.Err()
	if err != nil {
		return User{}, err
	}

	return data, nil
}

func (d *DB) GetUserByEmail(email string) (User, error) {
	db := d.db
	query :=
		`
  SELECT ID, Name, LastName, SurName, Email, PhoneNumber
  FROM User
  WHERE Email = ?
  `

	row := db.QueryRow(query, email)
	var data User
	err := row.Scan(
		&data.ID, &data.Name,
		&data.LastName, &data.SurName,
		&data.Email, &data.PhoneNumber,
	)
	if err != nil {
		return User{}, err
	}

	err = row.Err()
	if err != nil {
		return User{}, err
	}

	return data, nil
}

func (d *DB) GetUserTicketsByID(uid int) (Tickets, error) {
	db := d.db
	query :=
		`
  SELECT ID, DeparteLocID, ArriveLocID, Price, Airline,
  DepTime, DepDate
  FROM Ticket t
  JOIN UserTickets ut ON t.ID = ut.TicketID
  WHERE ut.UserID = ?;
  `

	rows, err := db.Query(query, uid)
	if err != nil {
		return Tickets{}, err
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
			return Tickets{}, err
		}
		data = append(data, ticket)
	}

	err = rows.Err()
	if err != nil {
		return Tickets{}, err
	}
	return data, nil
}

func (d *DB) GetUserFavoriteLocations(uid int) (Locations, error) {
	db := d.db
	query :=
		`
  SELECT l.LocName 
  FROM User u
  JOIN Favorites ul ON u.ID = ul.UserID
  JOIN Location l ON ul.LocationID = l.ID
  WHERE u.ID = ?
  `

	rows, err := db.Query(query, uid)
	if err != nil {
		return Locations{}, err
	}

	var data Locations
	for rows.Next() {
		var loc string
		err := rows.Scan(&loc)
		if err != nil {
			return Locations{}, err
		}
		data = append(data, loc)
	}
	err = rows.Err()
	if err != nil {
		return Locations{}, err
	}

	return data, nil
}

func (d *DB) GetTicketsByCitesAndDate(derlocid, arrlocid int, date1, date2 string, Isbusinss bool) (Tickets, error) {
	db := d.db
	query :=
		`
		SELECT ID, Airline, Price, DepDate, DepTime, TimeTaken 
		FROM Ticket
		WHERE DepDate >= STR_TO_DATE(?, '%Y-%m-%d')
		AND DepDate <= STR_TO_DATE(?, '%Y-%m-%d')
		AND DeparteLocID = ?
		AND ArriveLocID = ?
		AND IsBusiness = ? 
		ORDER BY Price ASC
  `

	rows, err := db.Query(
		query,
		date1, date2,
		derlocid, arrlocid,
		Isbusinss,
	)
	if err != nil {
		return Tickets{}, err
	}

	var data Tickets
	for rows.Next() {
		var ticket Ticket
		var tmp string
		err := rows.Scan(
			&ticket.ID,
			&ticket.Airline,
			&ticket.Price,
			&ticket.DepDate,
			&ticket.DepTime,
			&tmp,
		)

		arriveDate, _ := utils.GetArriveTime(ticket.DepDate, ticket.DepTime, tmp)
		ticket.ArriveDate = fmt.Sprintf("%v.%v.%v", arriveDate.Day(), arriveDate.Month(), arriveDate.Year())
		ticket.ArriveTime = fmt.Sprintf("%v:%v", arriveDate.Hour(), arriveDate.Minute())

		if err != nil {
			return Tickets{}, err
		}
		data = append(data, ticket)
	}

	err = rows.Err()
	if err != nil {
		return Tickets{}, err
	}
	return data, nil
}

func (d *DB) AddUser(user User, passwordHash string) (bool, error) {
	db := d.db
	query :=
		`
  INSERT INTO
  User(Name, Email, PasswordHash)
  VALUES
  (?, ?, ?)
  `

	result, err := db.Exec(query, user.Name, user.Email, passwordHash)
	if err != nil {
		return false, err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return false, nil
	}
	return true, nil
}

func (d *DB) AddTicketToFavorite(uid, locid int) (bool, error) {
	db := d.db
	query :=
		`
  INSERT INTO
  Favorites(UserID, LocationID)
  VALUES
  (?, ?)
  `

	result, err := db.Exec(query, uid, locid)
	if err != nil {
		return false, err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return false, nil
	}
	return true, nil
}

func (d *DB) UpdateUserInfo(uid int, user User) (bool, error) {
	db := d.db
	query :=
		`
  UPDATE User
  SET Name    = ?, LastName    = ?,
  SurName     = ?, Email       = ?,
  PhoneNumber = ?
  WHERE ID    = ?
  `

	result, err := db.Exec(
		query,
		user.Name, user.LastName, user.SurName,
		user.Email, user.PhoneNumber,
		uid,
	)

	if err != nil {
		return false, err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return false, nil
	}
	return true, nil
}

func (d *DB) DeleteUser(uid int) (bool, error) {
	db := d.db
	query :=
		`
  DELETE FROM User
  WHERE ID = ?
  `

	result, err := db.Exec(query, uid)
	if err != nil {
		return false, err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return false, nil
	}
	return true, nil

}

func (d *DB) DeleteTicketFromFavorite(uid, locid int) (bool, error) {
	db := d.db
	query :=
		`
  DELETE FROM Favorites
  WHERE UserID = ? AND LocationID = ?
  `

	result, err := db.Exec(query, uid, locid)
	if err != nil {
		return false, err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return false, nil
	}
	return true, nil
}

var _db *DB

func GetDB() *DB {
	if _db == nil {
		var db = DB{url: config.DBUrl}
		err := db.Connect()
		if err != nil {
			panic("[Error]: Cannot connect to database")
		}

		_db = &db
	}
	return _db
}
