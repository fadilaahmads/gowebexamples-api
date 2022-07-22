package repository

import (
	"fmt"
	"log"
	"time"
)

func GetSingleUser() {
	db, err := NewDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Query a single user
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}
