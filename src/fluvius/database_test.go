package main

import (
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {

	if len(db.Bookmarks(0)) != 0 {
		t.Fail()
	}
	db.SaveBookmark <- Bookmark{"test", "test", "test", "test", "test"}
	<-time.After(time.Duration(1 * time.Second))
	if len(db.Bookmarks(0)) != 1 {
		t.Fail()
	}
}
