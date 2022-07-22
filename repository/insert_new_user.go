package repository

import (
	"fmt"
	"log"
	"time"
)

func InsertNewUser() {
	db, err := NewDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Insert new user
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
