package repository

import (
	"fmt"
	"log"
	"time"
)

func GetAllUser() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	// Query all user
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		if err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal()
	}
	fmt.Printf("%#v", users)
}
