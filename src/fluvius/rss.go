package main

import (
	"fluvius/config"
	"github.com/luebken/rss"
	"log"
	"time"
)

func StartFetchingRss(configs []config.RSS) {
	for _, config := range configs {
		go fetchRss(config.Link, config.User)
	}
}

//TODO kind of wrong usage of feed library since it stores the items itself
func fetchRss(url string, user string) {
	log.Printf("Fetching %v", url)
	feed, err := rss.Fetch(url)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	log.Printf("Got feed %v. Current len(items): %v.\n", feed.Title, len(feed.Items))
	for _, item := range feed.Items {
		db.SaveBookmark <- Bookmark{
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
