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

const GET_ROOT_TEMPLATE = `
	<h1>Corpora</h1>
	<ul>
		{{range .Rows}}
			<li>{{.}}</li>
		{{end}}
	</ul>
`

func newRouter(dbConn *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleGetRoot(w, r, dbConn)
	}).Methods("GET")

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

func handleGetRoot(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	t, err := template.New("GET_ROOT_TEMPLATE").Parse(GET_ROOT_TEMPLATE)
	if err != nil {
		panic(err)
	}

	rows := db.FromCorpora(dbConn, "")

	data := struct {
		Rows []db.CorporaRow
	}{
		Rows: rows,
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
