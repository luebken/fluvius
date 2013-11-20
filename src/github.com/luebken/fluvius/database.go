package main

import (
	"log"
)

var db *database

//TODO better types
type Bookmark struct {
	Title   string
	Link    string
	Comment string
	User    string
	Feed    string
}

type database struct {
	save      chan Bookmark
	bookmarks map[string][]Bookmark
}

func init() {
	db = new(database)
	db.save = make(chan Bookmark)
	//out.res = make(chan bool)
	db.bookmarks = make(map[string][]Bookmark)
	go db.Run()
}

func (db *database) HotBookmarks() []Bookmark {
	return db.Bookmarks(1)
}

func (db *database) AllBookmarks() []Bookmark {
	return db.Bookmarks(0)
}

func (db *database) Bookmarks(size int) []Bookmark {
	result := []Bookmark{}
	for _, slice := range db.bookmarks {
		if len(slice) > size {
			merged := slice[0]
			merged.Comment = ""
			merged.User = ""
			for _, Bookmark := range slice {
				merged.Comment += Bookmark.User + ":" + Bookmark.Comment + ", "
				merged.User += Bookmark.User + ", "
			}
			merged.Comment = merged.Comment[:len(merged.Comment)-2]
			merged.User = merged.User[:len(merged.User)-2]
			result = append(result, merged)
		}
	}
	return result
}

func (db *database) Run() {
	var newBookmark Bookmark
	for {
		newBookmark = <-db.save
		slice, found := db.bookmarks[newBookmark.Link]
		if !found { //create slice with Bookmark
			db.bookmarks[newBookmark.Link] = []Bookmark{newBookmark}
		} else { //create or update Bookmark to a slice
			updated := false
			for index, value := range slice {
				//TODO: When do we want to update an Bookmark?
				if value.User == newBookmark.User {
					log.Printf("updating an Bookmark. %v \n", newBookmark)
					slice[index] = newBookmark
					updated = true
				}
			}
			if !updated {
				log.Printf("appending Bookmark %v to list \n", newBookmark)
				slice = append(slice, newBookmark)
				db.bookmarks[newBookmark.Link] = slice
			}
		}
	}
}
