package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertCorporaHasCorrectSchema(t *testing.T) {
	db := PrepareFakeDb()
	defer db.Close()

	assertCorporaHasCorrectSchema(db)
}

func TestFromCorpora(t *testing.T) {
	db := PrepareFakeDb()
	defer db.Close()

	rows := FromCorpora(db, "")
	assert.Equal(t, []CorporaRow{{1, "test"}}, rows)
}
