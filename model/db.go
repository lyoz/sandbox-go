package model

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:secret@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
}
