package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const HN_SEARCH = `http://api.thriftdb.com/api.hnsearch.com/items/_search?filter%5Bfields%5D%5Btype%5D=submission&q=`

type HNSearchResponse struct {
	Hits    int                      `json:"hits"`
	Time    float64                  `json:"time"`
	Results []HNSearchResponseResult `json:"results"`
}

type HNSearchResponseResult struct {
	Item  HNSearchResponseResultItem `json:"item"`
	Score float64                    `json:"score"`
}

type HNSearchResponseResultItem struct {
	Id     int `json:"id"`
	Points int `json:"points"`
}

func StartFetchingHN() {
	go runHN()
}

func runHN() {
	<-time.After(time.Duration(5 * time.Second))
	for {
		log.Println("Fetching HN")
		for _, y := range db.Bookmarks(0) {
			go fetchHN(y.Link)
			<-time.After(time.Duration(2 * time.Second))
		}
		<-time.After(time.Duration(60 * time.Second))
	}
}

//wraped in a function to make it mockable
var httpGet = func(bookmarkLink string) (*http.Response, error) {
	return http.Get(HN_SEARCH + bookmarkLink)
}

func fetchHN(bookmarkLink string) {
	r, err := httpGet(bookmarkLink)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	res := new(HNSearchResponse)
	dec.Decode(&res)
	if res.Hits > 0 {
		link := fmt.Sprintf("https://news.ycombinator.com/item?id=%d", res.Results[0].Item.Id)
		points := res.Results[0].Item.Points
		karma := Karma{BookmarkLink: bookmarkLink, Link: link, Points: points, Feed: "HN"}
		log.Printf("found karma: %v\n", karma)
		db.SaveKarma <- karma
	}

}
