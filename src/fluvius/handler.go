package main

import (
	"html/template"
	"net/http"
)

var tmpl = make(map[string]*template.Template)

func init() {
	tmpl = make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("base.html", "index.html"))
	tmpl["all.html"] = template.Must(template.ParseFiles("base.html", "all.html"))
}

type Page struct {
	Title     string
	PageItems []PageItem
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html")
	tmpl["index.html"].Execute(response, &Page{Title: "Fluvius – Hot", PageItems: db.Items(1)})
}

func AllHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html")
	tmpl["index.html"].Execute(response, &Page{Title: "Fluvius – All", PageItems: db.Items(0)})

}
