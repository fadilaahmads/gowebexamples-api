package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func DBConnect() {

	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	// Initialize the first connection to the database, to see if everything works correctly.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		username := "John Doe"
		password := "secret"
		createdAt := time.Now()

		// Inserts our data into the users table and returns with the result and a possible error.
		// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}

}
