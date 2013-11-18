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

	log.Printf("Got feed: %s\n")

	for _, item := range feed.Items {
		log.Printf("Appending %s\n", item.Title)
		log.Printf("\t%s\n", item.Link)
		log.Printf("\t%s\n\n", item.Content)

		AppendItems(Item{item.Title, item.Content, item.Link, "Oliver", feed.Title})
	}

	for {
		log.Println("updating feed")
		err = feed.Update()
		if err != nil {
			log.Printf("Error %v\n", err)
			//TODO handle error.
		}
		<-time.After(time.Duration(5 /*seconds*/ * 1e9))
	}
}
