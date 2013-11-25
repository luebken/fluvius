package main

import (
	cfg "fluvius/config"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = make(map[string]*template.Template)

func init() {
	fmt.Printf("assets dir: %v\n", cfg.AssetsDir())

	tmpl = make(map[string]*template.Template)

	tmpl["index.html"] = template.Must(template.ParseFiles(cfg.AssetsDir()+"base.html", cfg.AssetsDir()+"index.html"))
	tmpl["all.html"] = template.Must(template.ParseFiles(cfg.AssetsDir()+"base.html", cfg.AssetsDir()+"all.html"))
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
