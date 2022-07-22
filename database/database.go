package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

func DB() {

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

	{ // Insert new user
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

	{ // Query a single user
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

	{
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

	{
		// Delete user
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}

}
