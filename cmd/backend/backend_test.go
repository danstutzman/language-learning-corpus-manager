package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

func assertEqualHtml(t *testing.T, expected string, actual string) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(expected, actual, false)
	badDiffs := filterForBadDiffs(diffs)
	if len(badDiffs) > 0 {
		log.Printf("DiffPrettyText: " + dmp.DiffPrettyText(diffs))
		t.Fatal("Unexpected HTML")
	}
}

type Fixtures struct {
	dbConn *sql.DB
	router *mux.Router
	server *httptest.Server
}

func setupFixtures() *Fixtures {
	dbConn := db.PrepareFakeDb()
	router := newRouter(dbConn)
	server := httptest.NewServer(router)

	return &Fixtures{
		dbConn: dbConn,
		router: router,
		server: server,
	}
}

func teardownFixtures(fixtures *Fixtures) {
	fixtures.dbConn.Close()
	fixtures.server.Close()
}

func httpGet(t *testing.T, url string, expectedStatus int) string {
	resp, err := http.Get(url)
	assert.Nil(t, err)
	assert.Equal(t, expectedStatus, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func TestRouterGet404(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	httpGet(t, fixtures.server.URL+"/unknown", http.StatusNotFound)
}
