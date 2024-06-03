package config

import (
	"fmt"
	"os"
)

var User string = os.Getenv("FLYEASY_MYSQL_USER")
var Password string = os.Getenv("FLYEASY_MYSQL_PASSWORD")
var HostIP string = os.Getenv("FLYEASY_MYSQL_HOST")
var HostPort string = os.Getenv("FLYEASY_MYSQL_PORT")
var DBName string = os.Getenv("FLYEASY_MYSQL_DB")
var DBUrl string = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	User, Password, HostIP, HostPort, DBName)
