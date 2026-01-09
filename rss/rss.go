package rss

import (
	"github.com/charmbracelet/log"
	"github.com/mmcdole/gofeed"
)

// Returns true if the url is a valid feed.
func IsFeed(rssUrl string) bool {
	fp := gofeed.NewParser()
	_, err := fp.ParseURL(rssUrl)

	if err != nil {
		log.Error("the feed is not valid")
		return false
	}
	return true
}

func GetLatestPost(rssUrl string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rssUrl)
	log.Info(feed.Title)
}
