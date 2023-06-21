package data

import (
	"database/sql"
	"fmt"
	"log"
	"studying/web/config"
)

var DB *sql.DB

func DBConnection() *sql.DB {
	if DB == nil {
		var err error
		DB, err = sql.Open("mysql", config.DBConfig.FormatDSN())

		if err != nil {
			log.Fatal(err)
		}

		pingErr := DB.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
		}

		fmt.Println("[DB] Connected")
	}

	return DB
}
