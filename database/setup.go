package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	dsn := "root:pram123@tcp(mysql:3306)/db_go_api"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database : %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect database : %v", err)
	}
	DB = db
}
