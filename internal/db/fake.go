package db

import (
	"database/sql"
	"log"
)

func PrepareFakeDb() *sql.DB {
	connString := "file:temp.db?mode=memory"
	conn, err := sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatalf("Error from sql.Open: %s", err)
	}

	prepareFakeCorpora(conn)
	prepareFakeFiles(conn)

	return conn
}
