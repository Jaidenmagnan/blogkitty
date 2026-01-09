package rss

import (
	"github.com/Jaidenmagnan/blogkitty/db"
	"github.com/bwmarrin/discordgo"
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

// This is the cron job that will run and send out all feeds.
func Update(dg *discordgo.Session) {
	feeds, err := db.GetFeeds()
	if err != nil {
		log.Error("Could not get feeds.")
		return
	}

	for _, feed := range feeds {
		// Grab the latest post.
		latestPost, err := getLatestPost(feed.FeedURL)
		if err != nil {
			log.Error("Could not load latest post.", "url", feed.FeedURL)
		}

		// If the GUID is not the most recent one, don't send it again.
		if latestPost.GUID == feed.LatestPostGUID {
			log.Info("There are no new posts.", "feed", feed.FeedURL)
			continue
		}

		// Update the GUID field with the latest post.
		feed.LatestPostGUID = latestPost.GUID

		feed, err = db.UpdateFeed(feed)
		if err != nil {
			log.Error("Could not update the GUID field.")
		}

		// Send the message to the channel.
		// TODO: Remove channels that have been deleted.
		_, err = dg.ChannelMessageSend(feed.DiscordChannelID, latestPost.Link)
		if err != nil {
			log.Error("Could not send message.")
		}

		log.Info("New post sent.", "feed", feed.FeedURL)
	}
}

// Gets the latest post for a feed.
func getLatestPost(rssUrl string) (*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssUrl)
	if err != nil {
		log.Error("Could not parse feed.")
		return nil, err
	}

	if len(feed.Items) == 0 {
		log.Warn("No items in feed.")
		return nil, err
	}

	return feed.Items[0], nil
}
