package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/gowebexamples?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
