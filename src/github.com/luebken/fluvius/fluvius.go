// main package to start fluvius.
// fluvius a stream of information

package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// A fluvius item in the stream
type Item struct {
	Title   string
	Comment string
	Link    string
	User    string
	Feed    string
}

var items []Item = []Item{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler) //.Methods("GET")
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/img/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)

	log.Println("running server")
	startFetching()
	http.ListenAndServe(":8080", nil)
}

func AppendItem(item Item) {
	items = append(items, item)
}

func GetItems() []Item {
	return items
}
