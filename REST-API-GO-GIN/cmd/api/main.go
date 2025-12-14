package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}
	defer db.Close()
}
