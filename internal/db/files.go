package db

import (
	"database/sql"
	"fmt"
	"log"
)

type FilesRow struct {
	Id    int
	S3Key string
	Size  int
}

func prepareFakeFiles(db *sql.DB) {
	_, err := db.Exec(`
	  CREATE TABLE files (
      id     INTEGER PRIMARY KEY NOT NULL,
      s3_key TEXT NOT NULL,
      size   INTEGER NOT NULL
    );
    CREATE UNIQUE INDEX idx_files_s3_key ON files(s3_key);
    INSERT INTO files (s3_key, size) VALUES ('test.wav', 123);
  `)
	if err != nil {
		log.Fatal(err)
	}
}

func assertFilesHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, s3_key, size FROM files LIMIT 1"
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromFiles(db *sql.DB, whereLimit string) []FilesRow {
	query := "SELECT id, s3_key, size " +
		"FROM files " + whereLimit
	if LOG {
		log.Println(query)
	}

	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	rows := []FilesRow{}
	for rset.Next() {
		var row FilesRow
		err = rset.Scan(&row.Id, &row.S3Key, &row.Size)
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}

	err = rset.Err()
	if err != nil {
		panic(err)
	}

	return rows
}

func InsertFile(db *sql.DB, row FilesRow) FilesRow {
	query := fmt.Sprintf(`INSERT INTO files
    (s3_key, size)
    VALUES (%s, %d)`,
		EscapeString(row.S3Key),
		row.Size)
	if LOG {
		log.Println(query)
	}

	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	row.Id = int(id)

	return row
}
