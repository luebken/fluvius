package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Id int `json:"id"`
}

func startFetchingHN() {
	log.Println("start fetching hn")
	link := `http://techcrunch.com/2013/11/06/the-bitcoin-bubble/`

	go fetchHN(link)
}

func fetchHN(hnquery string) {

	hnsearch := `http://api.thriftdb.com/api.hnsearch.com/items/_search?filter%5Bfields%5D%5Btype%5D=submission&q=`

	r, err := http.Get(hnsearch + hnquery)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	res := new(HNSearchResponse)
	dec.Decode(&res)
	log.Printf("Got response %v\n", res)

	link := fmt.Sprintf("https://news.ycombinator.com/item?id=%d", res.Results[0].Item.Id)
	log.Printf("found %v", link)
}
