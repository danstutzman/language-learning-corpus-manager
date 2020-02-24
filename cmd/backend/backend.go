package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"log"
	"net/http"
	"os"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Specify a value for env key " + key)
		os.Exit(1)
	}
	return value
}

func main() {
	indexDownloader := index.NewS3IndexDownloader(
		getEnv("S3_REGION"),
		getEnv("S3_TOKEN"),
		getEnv("S3_SECRET"),
		getEnv("S3_BUCKET"),
		getEnv("TEMP_DIR"))
	dbConn, _, err := indexDownloader.Download()
	if err != nil {
		panic(err)
	}

	router := newRouter(dbConn, indexDownloader)

	log.Printf("Running web server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
