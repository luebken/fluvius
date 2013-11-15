package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	log.Println("running server")
	go Fetch()
	http.ListenAndServe(":8080", nil)
}
