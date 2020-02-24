package index

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
	"io"
	"log"
)

type FakeIndex struct {
	dbConn *sql.DB
}

func NewFakeIndex() *FakeIndex {
	return &FakeIndex{
		dbConn: db.PrepareFakeDb(),
	}
}

func (downloader *FakeIndex) Download() error {
	return nil
}

func (index *FakeIndex) ListCorpora() []db.CorporaRow {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	return db.FromCorpora(index.dbConn, "")
}

func (index *FakeIndex) ListFiles() []db.FilesRow {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	return db.FromFiles(index.dbConn, "")
}

func (index *FakeIndex) InsertFile(s3Key string, size int, reader io.Reader) {
	if index.dbConn == nil {
		log.Fatalf("Must call Download first")
	}

	db.InsertFile(index.dbConn, db.FilesRow{
		S3Key: s3Key,
		Size:  int(size),
	})
}
