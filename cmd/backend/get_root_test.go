package main

import (
	"net/http"
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

func TestRouterGetRoot(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	body := httpGet(t, fixtures.server.URL+"/", http.StatusOK)
	assertEqualHtml(t, EXPECTED_HTML_FOR_GET_ROOT, body)
}
