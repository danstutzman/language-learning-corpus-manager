package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter(index index.Index) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getRoot(w, r, index)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		getFiles(w, r, index)
	}).Methods("GET")

	r.HandleFunc("/files/new", func(w http.ResponseWriter, r *http.Request) {
		getFilesNew(w, r)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		postFiles(w, r, index)
	}).Methods("POST")

	return r
}
