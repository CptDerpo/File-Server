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

	sqlStmt := `CREATE TABLE IF NOT EXISTS User (
		id integer NOT NULL PRIMARY KEY, 
		username text, 
		password text, 
		admin integer
	);`

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

func addUserDB(username string, password string) error {
	user := DB.QueryRow(`SELECT username FROM User WHERE username = ?;`, username).Scan(&username)

	if user != sql.ErrNoRows {
		return errors.New("user already exists")
	}

	if len(password) < 8 {
		return errors.New("password is too short, must be at least 8 chars")
	}

	sqlStmt, _ := DB.Prepare("INSERT INTO User (username, password, admin) VALUES (?, ?, 0);")
	_, err := sqlStmt.Exec(username, password)
	return err
}

func removeUserDB(username string) error {
	sqlStmt, _ := DB.Prepare(`DELETE FROM User WHERE username=?`)
	_, err := sqlStmt.Exec(username)

	return err
}
