package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Items []Item
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Err %v", err)
	}
	t.Execute(response, &Page{Title: "Fluvius – Hot", Items: db.HotItems()})
}

func AllHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html")
	t, err := template.ParseFiles("all.html")
	if err != nil {
		log.Printf("Err %v", err)
	}
	t.Execute(response, &Page{Title: "Fluvius – All", Items: db.AllItems()})
}
