package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const EXPECTED_HTML_FOR_GET_ROOT = `
	<h1>Corpora</h1>
	<ul>
		<li>{1 test}</li>
	</ul>
`

func filterForBadDiffs(oldDiffs []diffmatchpatch.Diff) []diffmatchpatch.Diff {
	diffs := []diffmatchpatch.Diff{}
	for _, diff := range oldDiffs {
		if diff.Type != diffmatchpatch.DiffEqual &&
			strings.TrimSpace(diff.Text) != "" {
			diffs = append(diffs, diff)
		}
	}
	return diffs
}

func TestRouterGetRoot(t *testing.T) {
	db := db.PrepareFakeDb()
	defer db.Close()

	r := newRouter(db)
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(EXPECTED_HTML_FOR_GET_ROOT, string(b), false)
	badDiffs := filterForBadDiffs(diffs)
	if len(badDiffs) > 0 {
		log.Printf("DiffPrettyText: " + dmp.DiffPrettyText(diffs))
		t.Fatal("Unexpected HTML")
	}
}

func TestRouterGet405(t *testing.T) {
	db := db.PrepareFakeDb()
	defer db.Close()

	r := newRouter(db)
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/", "", nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "", string(b))
}
