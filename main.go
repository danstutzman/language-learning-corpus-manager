package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGetRoot).Methods("GET")
	return r
}

func main() {
	r := newRouter()

	log.Printf("Running web server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func handleGetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
