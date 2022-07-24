package repository

import (
	"fmt"
	"log"
)

func CreateUsersTable() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println("Table has been created!")
	defer db.Close()
}
