package main

import (
	"bitbucket.org/danstutzman/language-learning-corpus-manager/internal/db"
	"database/sql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
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

var getFilesNewTemplate = template.Must(template.New("getFilesNewTemplate").Parse(`<html>
	<form method='POST' action='/files' enctype="multipart/form-data">
		<h1>New File</h1>

		<p>
			<input type='file' name='file'>
		</p>

		<p>
			<input type='submit' value='Create File'>
		</p>
	</form>
</html>`))

func newRouter(dbConn *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getRoot(w, r, dbConn)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		getFiles(w, r, dbConn)
	}).Methods("GET")

	r.HandleFunc("/files/new", func(w http.ResponseWriter, r *http.Request) {
		getFilesNew(w, r)
	}).Methods("GET")

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		postFiles(w, r, dbConn)
	}).Methods("POST")

	return r
}

func main() {
	dbConn := db.InitDb(os.Getenv("DB_PATH"))

	r := newRouter(dbConn)

	log.Printf("Running web server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
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

func getFiles(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	files := db.FromFiles(dbConn, "")

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

func getFilesNew(w http.ResponseWriter, r *http.Request) {
	err := getFilesNewTemplate.Execute(w, struct{}{})
	if err != nil {
		panic(err)
	}
}

func postFiles(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	r.ParseMultipartForm(10 << 20) // Limit to 10MB file

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	db.InsertFile(dbConn, db.FilesRow{
		Filename: handler.Filename,
		Size:     int(handler.Size),
	})

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}
