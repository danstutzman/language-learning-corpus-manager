package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"bytes"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const EXPECTED_HTML_FOR_GET_ROOT = `<html>
  <p>
	  <a href='/files'>Files</a>
	</p>

	<h1>Corpora</h1>
	<ul>
		<li>{1 test}</li>
	</ul>
</html>`

const EXPECTED_HTML_FOR_GET_FILES = `<html>
	<h1>Files</h1>
	<ul>
		<li>{1 test.wav 123}</li>
	</ul>

	<p>
		<a href='/files/new'>New File</a>
	</p>
</html>`

const EXPECTED_HTML_FOR_GET_FILES_NEW = `<html>
	<form method='POST' action='/files' enctype="multipart/form-data">
		<h1>New File</h1>

		<p>
			<input type='file' name='file'>
		</p>

		<p>
			<input type='submit' value='Create File'>
		</p>
	</form>
</html>`

const EXPECTED_HTML_FOR_POST_FILES = `<html>
	<h1>Files</h1>
	<ul>
		<li>{1 test.wav 123}</li>
		<li>{2 transcript.txt 4}</li>
	</ul>

	<p>
		<a href='/files/new'>New File</a>
	</p>
</html>`

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

func TestRouterGetRoot(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_ROOT, body)
}

func TestRouterGet404(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	httpGet(t, fixtures.server.URL+"/unknown", http.StatusNotFound)
}

func TestRouterGetFiles(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/files", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_FILES, body)
}

func TestRouterGetFilesNew(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/files/new", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_FILES_NEW, body)
}

func TestRouterPostFiles(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	fileWriter, err := writer.CreateFormFile("file", "transcript.txt")
	assert.Nil(t, err)
	fileWriter.Write([]byte("test"))
	writer.Close()

	req, err := http.NewRequest("POST", fixtures.server.URL+"/files", &buffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := fixtures.server.Client()
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	assertEqualHtml(t, EXPECTED_HTML_FOR_POST_FILES, string(body))
}
