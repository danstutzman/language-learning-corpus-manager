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

	_, err = conn.Exec(`
		CREATE TABLE corpora (
			id   INTEGER PRIMARY KEY NOT NULL,
			name TEXT
		);
		CREATE UNIQUE INDEX idx_corpora_name ON corpora(name);
	  INSERT INTO corpora (name) VALUES ('test');

		CREATE TABLE files (
			id       INTEGER PRIMARY KEY NOT NULL,
			filename TEXT NOT NULL,
			size     INTEGER NOT NULL
		);
		CREATE UNIQUE INDEX idx_files_filename ON files(filename);
	  INSERT INTO files (filename, size) VALUES ('test.wav', 123);
	`)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
