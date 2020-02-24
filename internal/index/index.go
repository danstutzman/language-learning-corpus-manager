package index

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
)

type Index interface {
	Download() error
	ListCorpora() []db.CorporaRow
	ListFiles() []db.FilesRow
	InsertFile(filename string, size int)
}
