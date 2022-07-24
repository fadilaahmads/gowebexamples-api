package repository

import "log"

func DeleteSingleUser() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	// Delete user
	if _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
