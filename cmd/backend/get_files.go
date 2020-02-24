package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"html/template"
	"net/http"
)

var getFilesTemplate = template.Must(template.New("getFilesTempate").Parse(`<html>
	<h1>Files</h1>
	<ul>
		{{range .Files}}
			<li>{{.}}</li>
		{{end}}
	</ul>

	<p>
		<a href='/files/new'>New File</a>
	</p>
</html>`))

func getFiles(w http.ResponseWriter, r *http.Request, index index.Index) {
	files := index.ListFiles()

	data := struct {
		Files []db.FilesRow
	}{
		Files: files,
	}

	err := getFilesTemplate.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
