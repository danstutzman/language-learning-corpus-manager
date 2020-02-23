package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"log"
	"net/http"
	"os"
)

func main() {
	dbConn := db.InitDb(os.Getenv("DB_PATH"))

	router := newRouter(dbConn)

	log.Printf("Running web server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
