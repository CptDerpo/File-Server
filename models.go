package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createDB() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal("Error opening database")
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS User (id integer not null primary key, username text, password text, admin integer);
	CREATE TABLE IF NOT EXISTS File (id integer not null primary key, filename text, path string, filtetype text, date datetime, size int);
	`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
