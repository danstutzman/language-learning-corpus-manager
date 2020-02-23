package db

import (
	"database/sql"
	"log"
)

type CorporaRow struct {
	Id   int
	Name string
}

func assertCorporaHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, name FROM corpora LIMIT 1"
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromCorpora(db *sql.DB, whereLimit string) []CorporaRow {
	query := "SELECT id, name " +
		"FROM corpora " + whereLimit
	if LOG {
		log.Println(query)
	}

	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	rows := []CorporaRow{}
	for rset.Next() {
		var row CorporaRow
		err = rset.Scan(&row.Id, &row.Name)
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
