package index

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
)

type FakeIndexDownloader struct {
}

func NewFakeIndexDownloader() *FakeIndexDownloader {
	return &FakeIndexDownloader{}
}

func (downloader *FakeIndexDownloader) Download() (*sql.DB, string, error) {
	return db.PrepareFakeDb(), "", nil
}
