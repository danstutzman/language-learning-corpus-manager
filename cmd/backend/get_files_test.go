package main

import (
	"net/http"
	"testing"
)

const EXPECTED_HTML_FOR_GET_FILES = `<html>
	<h1>Files</h1>
	<ul>
		<li>{1 test.wav 123}</li>
	</ul>

	<p>
		<a href='/files/new'>New File</a>
	</p>
</html>`

func TestRouterGetFiles(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/files", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_FILES, body)
}
