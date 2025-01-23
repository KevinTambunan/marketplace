package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:@/marketplace?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	}
	log.Println("Database connection closed")
}
