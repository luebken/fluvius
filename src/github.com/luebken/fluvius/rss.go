package main

import (
	rss "github.com/SlyMarbo/rss"
	"log"
	"time"
)

func startFetchingRss() {
	log.Println("start fetching oli")
	go fetchRss("https://feeds.pinboard.in/atom/u:othylmann/", "Oliver")
	log.Println("start fetching matthias")
	go fetchRss("https://feeds.pinboard.in/atom/u:luebken/", "Matthias")
}

//TODO kind of wrong usage of feed library since it stores the items itself
func fetchRss(url string, user string) {
	feed, err := rss.Fetch(url)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	log.Printf("Got feed %v. Current len(items): %v.\n", feed.Title, len(feed.Items))
	for _, item := range feed.Items {
		db.saveBookmark <- Bookmark{
			Title:   item.Title,
			Comment: item.Content,
			Link:    item.Link,
			User:    user,
			Feed:    feed.Title}
	}

	for {
		log.Printf("Updating feed %v. Current len(items): %v.\n", feed.Title, len(feed.Items))
		err = feed.Update()
		if err != nil {
			log.Printf("Error %v\n", err)
		}
		log.Printf("Updated feed %v. Now len(items): %v.\n", feed.Title, len(feed.Items))
		/* TODO need to check wether items are already
		for _, item := range feed.Items {
			AppendItem(Item{
				Title:   item.Title,
				Comment: item.Content,
				Link:    item.Link,
				User:    user,
				Feed:    feed.Title})
		}
		*/

		<-time.After(time.Duration(60 * time.Second))
	}
}
