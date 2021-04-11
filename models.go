package main

import (
	"database/sql"
	"errors"
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

//opens a db but needs to be manually closed afterwards
func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal("Error opening database")
	}
	return db
}

func addUserDB(username string, password string, db *sql.DB) error {
	user := db.QueryRow(`SELECT username FROM User WHERE username = ?`, username).Scan(&username)

	if user != sql.ErrNoRows {
		return errors.New("user already exists")
	}

	if len(password) < 8 {
		return errors.New("password is too short, must be at least 8 chars")
	}

	sqlStmt, _ := db.Prepare("INSERT INTO User (username, password, admin) VALUES (?, ?, 0)")
	sqlStmt.Exec(username, password)
	return nil
}

//implement exception
func removeUserDB(username string, db *sql.DB) error {
	sqlStmt, _ := db.Prepare(`DELETE FROM User WHERE username=?`)
	sqlStmt.Exec(username)
	return nil
}
