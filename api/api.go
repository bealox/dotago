package api

import (
	"os"

	//Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = os.Getenv("MySQLUser")
	password = os.Getenv("MySQLPassword")
)
