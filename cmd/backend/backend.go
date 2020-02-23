package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"net/http"
	"os"
	"path"
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
	s3Region := getEnv("S3_REGION")
	s3Token := getEnv("S3_TOKEN")
	s3Secret := getEnv("S3_SECRET")
	session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Credentials: credentials.NewStaticCredentials(s3Token, s3Secret, ""),
	}))

	tempDir := getEnv("TEMP_DIR")
	file, err := os.Create(path.Join(tempDir, "index.sqlite3"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s3Bucket := getEnv("S3_BUCKET")
	downloader := s3manager.NewDownloader(session)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String("index.sqlite3"),
		})
	if err != nil {
		panic(err)
	}
	log.Printf("Wrote index.sqlite3")

	//////
	dbConn := db.InitDb(getEnv("DB_PATH"))

	router := newRouter(dbConn)

	log.Printf("Running web server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
