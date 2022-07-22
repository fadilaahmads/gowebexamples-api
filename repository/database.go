package repository

import (
	"database/sql"
	"fmt"
	"log"
)

func NewDB() *sql.DB {
	var (
		host     = "127.0.0.1"
		port     = "3306"
		username = "root"
		password = ""
		dbname   = "gowebexamples"
	)

	psqlInfo := fmt.Sprintf("host= %s port= %s user= %s "+"password= %s dbname= %s", host, port, username, password, dbname)

	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal()
	}

	return db
}
