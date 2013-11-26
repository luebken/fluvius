package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bm1 = Bookmark{"Google", "http://google.de", "test", "test", "test"}
var bm2 = Bookmark{"Bing", "http://bing.de", "test", "test", "test"}
var ka1 = Karma{"http://google.de", "https://news.ycombinator.com/item?id=1", 1, "test"}
var ka2 = Karma{"http://bing.de", "http://www.reddit.com/r/technology/comments/1rffrf", 1, "test"}

var _ = Describe("Database", func() {
	Context("Bookmarks", func() {
		It("should be empty", func() {
			Expect(len(db.Bookmarks(0))).To(Equal(0))
		})
		It("should store one bookmark", func(done Done) {
			listener := make(chan Bookmark)
			db.addEventListener(listener)
			db.SaveBookmark <- bm1
			Expect(<-listener).To(Equal(bm1))
			Expect(len(db.Bookmarks(0))).To(Equal(1))
			db.removeEventListener(listener)
			close(done)
		})
		It("should store a second bookmark", func(done Done) {
			listener := make(chan Bookmark)
			db.addEventListener(listener)
			db.SaveBookmark <- bm2
			Expect(<-listener).To(Equal(bm2))
			Expect(len(db.Bookmarks(0))).To(Equal(2))
			db.removeEventListener(listener)
			close(done)
		})
	})

	Context("Karma", func() {
		It("should be empty", func() {
			Expect(len(db.karmas)).To(Equal(0))
		})
		It("should store one karma", func(done Done) {
			Expect(len(db.Bookmarks(0))).To(Equal(2))
			listener := make(chan Karma)
			db.addEventListener(listener)
			db.SaveKarma <- ka1
			Expect(<-listener).To(Equal(ka1))
			Expect(len(db.karmas)).To(Equal(1))
			db.removeEventListener(listener)
			close(done)
		})
		It("should store a second karma", func(done Done) {
			Expect(len(db.Bookmarks(0))).To(Equal(2))
			Expect(len(db.karmas)).To(Equal(1))

			listener := make(chan Karma)
			db.addEventListener(listener)
			db.SaveKarma <- ka2
			Expect(<-listener).To(Equal(ka2))
			Expect(len(db.karmas)).To(Equal(2))
			db.removeEventListener(listener)
			close(done)
		})
	})

})
