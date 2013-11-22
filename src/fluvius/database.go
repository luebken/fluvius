package main

import (
	"log"
)

var db *Database

//Bookmarks with comments
type Bookmark struct {
	Title   string
	Link    string
	Comment string
	User    string
	Feed    string
}

//Other references to bookmarks like HN
type Karma struct {
	BookmarkLink string
	Link         string
	Points       int
	Feed         string
}

//The struct to be displayed on a webpage
type PageItem struct {
	Title    string
	Link     string
	Comments string
	Users    string
	Karmas   []Karma
}

type Database struct {
	SaveBookmark chan Bookmark
	SaveKarma    chan Karma
	bookmarks    map[string][]Bookmark
	karmas       map[string][]Karma
}

func init() {
	db = new(Database)
	db.SaveBookmark = make(chan Bookmark)
	db.bookmarks = make(map[string][]Bookmark)
	db.SaveKarma = make(chan Karma)
	db.karmas = make(map[string][]Karma)
	go db.runBookmarks()
	go db.runKarmas()
}

func (db *Database) Items(bookmarksThreshold int) []PageItem {

	log.Printf("getting Items. karmas: %v\n", db.karmas)
	result := []PageItem{}
	bookmarks := db.Bookmarks(bookmarksThreshold)
	for _, bookmark := range bookmarks {
		item := PageItem{}
		item.Title = bookmark.Title
		item.Link = bookmark.Link
		item.Comments = bookmark.Comment
		item.Users = bookmark.User
		item.Karmas = db.karmas[item.Link]
		result = append(result, item)
	}
	return result
}

//TODO: not clean to have the same type for a single bookmark and a merged bookmarks
func (db *Database) Bookmarks(bookmarksThreshold int) []Bookmark {
	result := []Bookmark{}
	for _, slice := range db.bookmarks {
		if len(slice) > bookmarksThreshold {
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

// private function
func (db *Database) runBookmarks() {
	var newBookmark Bookmark
	for {
		newBookmark = <-db.SaveBookmark
		slice, found := db.bookmarks[newBookmark.Link]
		if !found { //create slice with Bookmark
			db.bookmarks[newBookmark.Link] = []Bookmark{newBookmark}
		} else { //create or update Bookmark to a slice
			updated := false
			for index, value := range slice {
				//TODO: When do we want to update an Bookmark?
				if value.User == newBookmark.User {
					log.Printf("updating a bookmark. %v \n", newBookmark)
					slice[index] = newBookmark
					updated = true
				}
			}
			if !updated {
				log.Printf("appending a bookmark %v to list \n", newBookmark)
				slice = append(slice, newBookmark)
				db.bookmarks[newBookmark.Link] = slice
			}
		}
	}
}

// private function
func (db *Database) runKarmas() {
	var newKarma Karma
	for {
		newKarma = <-db.SaveKarma
		slice, found := db.karmas[newKarma.BookmarkLink]
		if !found { //create slice
			db.karmas[newKarma.BookmarkLink] = []Karma{newKarma}
		} else { //create or update Bookmark to a slice
			updated := false
			for index, value := range slice {
				//TODO: When do we want to update a Karma?
				if value.Link == newKarma.Link {
					slice[index] = newKarma
					updated = true
				}
			}
			if !updated {
				slice = append(slice, newKarma)
				db.karmas[newKarma.BookmarkLink] = slice
			}
		}
	}
}
