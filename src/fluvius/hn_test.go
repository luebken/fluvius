package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"log"
	"net/http"
)

var buffer = bytes.NewBufferString("your string")

type DummyReaderWriter struct {
}

func (d DummyReaderWriter) Read(p []byte) (n int, err error) {
	return buffer.Read(p)
}
func (d DummyReaderWriter) Close() error {
	return nil
}

var _ = Describe("Hn", func() {
	Context("fetch HN", func() {

		It("should be empty", func() {

			var bookmarkLink = "www.google.de"
			httpGet = func(bookmarkLink string) (*http.Response, error) {
				var response = new(http.Response)
				response.Body = new(DummyReaderWriter)
				return response, nil
			}

			fetchHN(bookmarkLink)
			log.Printf("%s", httpGet)
			//TODO real assertion
			Expect(0).To(Equal(0))
		})
	})
})
