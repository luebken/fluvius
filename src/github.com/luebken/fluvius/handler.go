package fluvius

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
	log.Println("requested /")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Err %v", err)
	}
	t.Execute(response, &Page{Title: "Fluvius ––– Stream", Items: GetItems()})
}
