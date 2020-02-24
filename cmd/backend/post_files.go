package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"net/http"
)

func postFiles(w http.ResponseWriter, r *http.Request, index index.Index) {
	r.ParseMultipartForm(10 << 20) // Limit to 10MB file

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	index.InsertFile(handler.Filename, int(handler.Size))

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}
