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

func (db *database) Items() []Item {
	result := []Item{}
	for _, slice := range db.store {
		for _, value := range slice {
			result = append(result, value)
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
					log.Printf("updating an item \n")
					slice[index] = newItem
					updated = true
				}
			}
			if !updated {
				log.Println("appending new item to list")
				slice = append(slice, newItem)
				db.store[newItem.Link] = slice
			}
		}
	}
}
