package main

import (
	"log"
)

var db *database

//TODO better types
type Item struct {
	Title   string
	Link    string
	Comment string
	User    string
	Feed    string
}

type database struct {
	save  chan Item
	store map[string][]Item
}

func init() {
	db = new(database)
	db.save = make(chan Item)
	//out.res = make(chan bool)
	db.store = make(map[string][]Item)
	go db.Run()
}

func (db *database) HotItems() []Item {
	return db.Items(1)
}

func (db *database) AllItems() []Item {
	return db.Items(0)
}

func (db *database) Items(size int) []Item {
	result := []Item{}
	for _, slice := range db.store {
		if len(slice) > size {
			merged := slice[0]
			merged.Comment = ""
			merged.User = ""
			for _, item := range slice {
				merged.Comment += item.User + ":" + item.Comment + ", "
				merged.User += item.User + ", "
			}
			merged.Comment = merged.Comment[:len(merged.Comment)-2]
			merged.User = merged.User[:len(merged.User)-2]
			result = append(result, merged)
		}
	}
	return result
}

func (db *database) Run() {
	var newItem Item
	for {
		newItem = <-db.save
		slice, found := db.store[newItem.Link]
		if !found { //create slice with item
			db.store[newItem.Link] = []Item{newItem}
		} else { //create or update item to a slice
			updated := false
			for index, value := range slice {
				//TODO: When do we want to update an item?
				if value.User == newItem.User {
					log.Printf("updating an item. %v \n", newItem)
					slice[index] = newItem
					updated = true
				}
			}
			if !updated {
				log.Printf("appending item %v to list \n", newItem)
				slice = append(slice, newItem)
				db.store[newItem.Link] = slice
			}
		}
	}
}
