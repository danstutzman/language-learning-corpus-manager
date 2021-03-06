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
		<li>{2 test/transcript.txt 4}</li>
	</ul>

	<p>
		<a href='/files/new'>New File</a>
	</p>
</html>`

func TestRouterPostFiles(t *testing.T) {
	fixtures := setupFixtures()
	defer teardownFixtures(fixtures)

	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)

	fileWriter, err := multipartWriter.CreateFormFile("file", "transcript.txt")
	assert.Nil(t, err)
	fileWriter.Write([]byte("test"))

	fieldWriter, err := multipartWriter.CreateFormField("corpus_name")
	assert.Nil(t, err)
	fieldWriter.Write([]byte("test"))

	multipartWriter.Close()

	req, err := http.NewRequest("POST", fixtures.server.URL+"/files", &buffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := fixtures.server.Client()
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	assertEqualHtml(t, EXPECTED_HTML_FOR_POST_FILES, string(body))
}
