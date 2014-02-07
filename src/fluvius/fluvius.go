// main package to start fluvius.
// fluvius a stream of information

package main

import (
	"github.com/gorilla/mux"
	"log"
	//"time"
	"fluvius/config"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/all.html", AllHandler).Methods("GET")
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/img/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)

	log.Println("running server")
	StartFetchingRss(config.GetRssConfig())
	StartFetchingHN()
	StartFetchingTwitter(
		config.GetTwitterConsumerKey(),
		config.GetTwitterConsumerSecret(),
		config.GetTwitterAccessToken(),
		config.GetTwitterAccessTokenSecret())
	http.ListenAndServe(":8080", nil)
}
