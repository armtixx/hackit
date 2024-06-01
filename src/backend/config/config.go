package config

import (
	"fmt"
	"os"
)

var User string = os.Getenv("USNAME")
var Password string = os.Getenv("PASWD")
var HostIP string = os.Getenv("HOST_IP")
var HostPort string = os.Getenv("HOST_PORT")
var DBName string = os.Getenv("DB_NAME")
var DBURL string = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	User, Password, HostIP, HostPort, DBName)
