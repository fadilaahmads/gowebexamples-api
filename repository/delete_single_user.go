package repository

import "log"

func DeleteSingleUser(id string) {
	db, err := NewDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Delete user
	if _, err := db.Exec(`DELETE FROM users WHERE id = ?`, id); err != nil {
		log.Fatal(err)
	}

}
