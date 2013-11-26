package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database", func() {
	Context("Empty Database", func() {
		It("should be empty", func() {
			Expect(len(db.Bookmarks(0))).To(Equal(0))
		})
		It("should store one bookmark", func(done Done) {

			bm := Bookmark{"test", "test", "test", "test", "test"}
			listener := make(chan Bookmark)
			db.addBookmarkEventListener(listener)
			db.SaveBookmark <- bm
			Expect(<-listener).To(Equal(bm))
			Expect(len(db.Bookmarks(0))).To(Equal(1))
			close(done)
		})
	})

})
