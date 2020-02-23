package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const LOG = true

func InitDb(dbPath string) *sql.DB {
	// Set mode=rw so it doesn't create database if file doesn't exist
	connString := "file:" + dbPath + "?mode=rw"
	dbConn, err := sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatalf("Error from sql.Open: %s", err)
	}

	assertCorporaHasCorrectSchema(dbConn)

	return dbConn
}
