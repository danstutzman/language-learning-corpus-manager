package index

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type S3Index struct {
	s3Bucket  string
	s3Session *session.Session
	dbPath    string
	dbConn    *sql.DB
}

func NewS3Index(s3Region, s3Token, s3Secret, s3Bucket,
	tempDir string) *S3Index {
	s3Session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Credentials: credentials.NewStaticCredentials(s3Token, s3Secret, ""),
	}))

	return &S3Index{
		s3Bucket:  s3Bucket,
		s3Session: s3Session,
		dbPath:    path.Join(tempDir, "index.sqlite3"),
		dbConn:    nil,
	}
}

func (index *S3Index) Download() error {
	log.Printf("Downloading " + index.dbPath + " from S3...")

	service := s3.New(index.s3Session)
	getOutput, err := service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(index.s3Bucket),
		Key:    aws.String("index.sqlite3"),
	})
	if err != nil {
		return err
	}

	dbFile, err := os.Create(index.dbPath)
	if err != nil {
		return err
	}
	defer dbFile.Close()

	bytes, err := ioutil.ReadAll(getOutput.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(index.dbPath, bytes, 0644)
	if err != nil {
		return err
	}
	log.Printf("Wrote " + index.dbPath)

	index.dbConn = db.InitDb(index.dbPath)
	return nil
}

func (index *S3Index) ListCorpora() []db.CorporaRow {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	return db.FromCorpora(index.dbConn, "")
}

func (index *S3Index) ListFiles() []db.FilesRow {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	return db.FromFiles(index.dbConn, "")
}

func (index *S3Index) InsertFile(filename string, size int) {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	db.InsertFile(index.dbConn, db.FilesRow{
		Filename: filename,
		Size:     int(size),
	})
}
