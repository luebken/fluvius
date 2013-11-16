package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var items []string = []string{"item 1", "item 2"}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	http.Handle("/", r)

	log.Println("running server")
	go Fetch()
	http.ListenAndServe(":8080", nil)
}

func AppendItems(newItem string) {
	items = append(items, newItem)
}

func GetItems() []string {
	return items
}
