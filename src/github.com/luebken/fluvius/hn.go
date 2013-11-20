package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func startFetchingHN() {
	go runHN()
}

func runHN() {
	<-time.After(time.Duration(5 * time.Second))

	log.Println("start fetching hn")
	//link := `http://techcrunch.com/2013/11/06/the-bitcoin-bubble/`

	for {
		log.Println("Fetching HN")
		for _, y := range db.HotItems() {
			go fetchHN(y.Link)

		}
		<-time.After(time.Duration(60 * time.Second))
	}
}

func fetchHN(hnquery string) {
	log.Printf("HN query for %v", hnquery)
	hnsearch := `http://api.thriftdb.com/api.hnsearch.com/items/_search?filter%5Bfields%5D%5Btype%5D=submission&q=`

	r, err := http.Get(hnsearch + hnquery)
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
		log.Printf("found Points:%v Link:%v", points, link)
	}

}
