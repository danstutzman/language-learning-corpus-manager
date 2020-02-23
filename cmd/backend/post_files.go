package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
	"net/http"
)

func postFiles(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	r.ParseMultipartForm(10 << 20) // Limit to 10MB file

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	db.InsertFile(dbConn, db.FilesRow{
		Filename: handler.Filename,
		Size:     int(handler.Size),
	})

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}
