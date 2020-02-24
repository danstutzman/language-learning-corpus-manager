package main

import (
	"net/http"
	"testing"
)

const EXPECTED_HTML_FOR_GET_FILES_NEW = `<html>
	<form method='POST' action='/files' enctype="multipart/form-data">
		<h1>New File</h1>

		<p>
			<label>Corpus name</label><br>
			<input name='corpus_name' value='spintx'>
		</p>

		<p>
			<input type='file' name='file'>
		</p>

		<p>
			<input type='submit' value='Create File'>
		</p>
	</form>
</html>`

func TestRouterGetFilesNew(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/files/new", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_FILES_NEW, body)
}
