package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Items []string
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html")
	log.Println("ROOT")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Err %v", err)
	}
	t.Execute(response, &Page{Title: "a title", Items: GetItems()})
}
