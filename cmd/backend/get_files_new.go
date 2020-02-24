package main

import (
	"html/template"
	"net/http"
)

var getFilesNewTemplate = template.Must(template.New("getFilesNewTemplate").Parse(`<html>
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
</html>`))

func getFilesNew(w http.ResponseWriter, r *http.Request) {
	err := getFilesNewTemplate.Execute(w, struct{}{})
	if err != nil {
		panic(err)
	}
}
