package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

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
