package main

import (
	rss "github.com/SlyMarbo/rss"
	"log"
	"time"
)

func Fetch() {
	feed, err := rss.Fetch("https://feeds.pinboard.in/atom/u:othylmann/")
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	for {
		log.Println("updating feed")
		err = feed.Update()
		if err != nil {
			log.Printf("Error %v\n", err)
			// handle error.
		}
		<-time.After(time.Duration(1 /*seconds*/ * 1e9))
	}
}
