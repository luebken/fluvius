package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

func runTwitter(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	fmt.Println("getSearch 1")
	v := make(map[string][]string)
	v["screen_name"] = []string{"luebken"}
	v["count"] = []string{"20"}
	searchResult, _ := api.GetHomeTimeline(v)

	for _, tweet := range searchResult {
		if len(tweet.Entities.Urls) > 0 {
			fmt.Println(tweet.Entities.Urls[0].Expanded_url)
		}
	}
}
