package config

import (
	"go/build"
	"log"
	"os"
	"strings"
)

type RSS struct {
	Link string
	User string
}

var rssConfig []RSS

var assetsDir string

func init() {
	log.Println("Initializing config")
	s := os.Getenv("FLUVIUS_RSS_FEEDS")
	for _, s2 := range strings.Split(s, "||") {
		user := strings.TrimSpace(strings.Split(s2, "|")[0])
		link := strings.TrimSpace(strings.Split(s2, "|")[1])
		rssConfig = append(rssConfig, RSS{user, link})
	}
	if len(rssConfig) == 0 {
		panic("couldn't find rss config see readme.md")
	}

}

func GetRssConfig() []RSS {
	return rssConfig
}

func GetTwitterConsumerKey() string {
	return os.Getenv("TWITTER_CONSUMER_KEY")
}

func GetTwitterConsumerSecret() string {
	return os.Getenv("TWITTER_CONSUMER_SECRET")
}

func GetTwitterAccessToken() string {
	return os.Getenv("TWITTER_ACCESS_TOKEN")
}

func GetTwitterAccessTokenSecret() string {
	return os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
}

func AssetsDir() string {
	if assetsDir == "" {
		p, err := build.Default.Import("fluvius", "", build.FindOnly)
		if err != nil {
			log.Panicf("Can't load files %v\n", err)
		}
		assetsDir = p.Dir + "/assets/"
		log.Printf("Loading templates from:%v\n", assetsDir)
	}
	return assetsDir
}
