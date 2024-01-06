package config

import (
	"fmt"
	"log"
)

func CheckIDdb(tabel string, id int) bool {
	db, err := MySQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE id = %v", tabel, id)
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		log.Println("Error querying the database:", err)
		return false
	}

	return count > 0

}
