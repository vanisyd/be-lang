package config

import "github.com/go-sql-driver/mysql"

var DBConfig = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	Net:                  "tcp",
	Addr:                 "127.0.0.1:3306",
	DBName:               "study",
	AllowNativePasswords: true,
}
