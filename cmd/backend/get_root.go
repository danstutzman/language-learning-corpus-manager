package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/index"
	"database/sql"
	"html/template"
	"net/http"
)

var getRootTemplate = template.Must(template.New("getRootTempate").Parse(`<html>
  <p>
	  <a href='/files'>Files</a>
	</p>

	<h1>Corpora</h1>
	<ul>
		{{range .Corpora}}
			<li>{{.}}</li>
		{{end}}
	</ul>
</html>`))

func getRoot(w http.ResponseWriter, r *http.Request, dbConn *sql.DB,
	downloader index.IndexDownloader) {

	corpora := db.FromCorpora(dbConn, "")

	data := struct {
		Corpora []db.CorporaRow
	}{
		Corpora: corpora,
	}

	err := getRootTemplate.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
