package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bm1 = Bookmark{"Google", "http://google.de", "test", "test", "test"}
var bm2 = Bookmark{"Bing", "http://bing.de", "test", "test", "test"}

var _ = Describe("Database", func() {
	Context("Bookmarks", func() {
		It("should be empty", func() {
			Expect(len(db.Bookmarks(0))).To(Equal(0))
		})
		It("should store one bookmark", func(done Done) {
			listener := make(chan Bookmark)
			db.addBookmarkEventListener(listener)
			db.SaveBookmark <- bm1
			Expect(<-listener).To(Equal(bm1))
			Expect(len(db.Bookmarks(0))).To(Equal(1))
			db.removeBookmarkEventListener(listener)
			close(done)
		})
		It("should store a second bookmark", func(done Done) {
			listener := make(chan Bookmark)
			db.addBookmarkEventListener(listener)
			db.SaveBookmark <- bm2
			Expect(<-listener).To(Equal(bm2))
			Expect(len(db.Bookmarks(0))).To(Equal(2))
			db.removeBookmarkEventListener(listener)
			close(done)
		})
	})
	/*
		Context("Karma", func() {
			It("should be empty", func() {
				Expect(len(db.karmas)).To(Equal(0))
			})
			It("should store one karma", func(done Done) {
				bm = Bookmark{"test", "test", "test", "test", "test"}
				listener := make(chan Bookmark)
				db.addBookmarkEventListener(listener)
				db.SaveBookmark <- bm
				Expect(<-listener).To(Equal(bm))
				Expect(len(db.Bookmarks(0))).To(Equal(1))
				close(done)
			})
		})
	*/

})
