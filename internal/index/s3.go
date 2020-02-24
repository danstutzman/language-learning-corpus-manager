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

type IndexDownloader interface {
	Download() (*sql.DB, string, error)
}

type S3IndexDownloader struct {
	s3Bucket  string
	s3Session *session.Session
	dbPath    string
}

func NewS3IndexDownloader(s3Region, s3Token, s3Secret, s3Bucket,
	tempDir string) *S3IndexDownloader {
	s3Session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Credentials: credentials.NewStaticCredentials(s3Token, s3Secret, ""),
	}))

	downloader := &S3IndexDownloader{
		s3Bucket:  s3Bucket,
		s3Session: s3Session,
		dbPath:    path.Join(tempDir, "index.sqlite3"),
	}

	return downloader
}

// 2nd return is the new etag
func (downloader *S3IndexDownloader) Download() (*sql.DB, string, error) {
	log.Printf("Downloading " + downloader.dbPath + " from S3...")

	service := s3.New(downloader.s3Session)
	getOutput, err := service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(downloader.s3Bucket),
		Key:    aws.String("index.sqlite3"),
	})
	if err != nil {
		return nil, "", err
	}

	dbFile, err := os.Create(downloader.dbPath)
	if err != nil {
		return nil, "", err
	}
	defer dbFile.Close()

	bytes, err := ioutil.ReadAll(getOutput.Body)
	if err != nil {
		return nil, "", err
	}

	err = ioutil.WriteFile(downloader.dbPath, bytes, 0644)
	if err != nil {
		return nil, "", err
	}
	log.Printf("Wrote " + downloader.dbPath)

	return db.InitDb(downloader.dbPath), *getOutput.ETag, nil
}
