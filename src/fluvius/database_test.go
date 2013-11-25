package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Database", func() {
	Context("Empty Database", func() {
		It("should be empty", func() {
			Expect(len(db.Bookmarks(0))).To(Equal(0))
		})
		It("should store one bookmakr", func() {
			db.SaveBookmark <- Bookmark{"test", "test", "test", "test", "test"}
			<-time.After(time.Duration(1 * time.Microsecond)) //TODO remove
			Expect(len(db.Bookmarks(0))).To(Equal(1))
		})
	})

})
