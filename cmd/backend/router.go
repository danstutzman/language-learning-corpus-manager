package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter(dbConn *sql.DB,
	indexDownloader index.IndexDownloader) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getRoot(w, r, dbConn, indexDownloader)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		getFiles(w, r, dbConn)
	}).Methods("GET")

	r.HandleFunc("/files/new", func(w http.ResponseWriter, r *http.Request) {
		getFilesNew(w, r)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		postFiles(w, r, dbConn)
	}).Methods("POST")

	return r
}
