// This package initializes the db.
package db

import (
	"database/sql"
	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

// The global DB connection.
var DB *sql.DB

// Connects to the database.
func Connect() error {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Info("Database connected successfully")
	return nil
}

type Feed struct {
	ID               int64
	LatestPostGUID   string
	DiscordChannelID string
	FeedURL          string
}

// Add a feed to the database and populate the ID field.
func InsertFeed(feed Feed) (Feed, error) {
	query := "INSERT INTO feeds (feed_url, discord_channel_id, latest_post_guid) VALUES (?, ?, ?)"

	result, err := DB.Exec(query, feed.FeedURL, feed.DiscordChannelID, feed.LatestPostGUID)
	if err != nil {
		log.Error("error inserting feed")
		return Feed{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error("failed getting the latest inserted ID")
		return Feed{}, err
	}

	feed.ID = id
	return feed, nil
}

// Update the GUID of a feed to match the latest post.
func UpdateFeed(feed Feed) (Feed, error) {
	query := "UPDATE feeds SET latest_post_guid = ?, discord_channel_id = ?, feed_url = ? WHERE id = ?"

	_, err := DB.Exec(query, feed.LatestPostGUID, feed.DiscordChannelID, feed.FeedURL, feed.ID)
	if err != nil {
		log.Error("error updating the guid field")
		return Feed{}, err
	}

	return feed, nil

}

// Get all feeds.
func GetFeeds() ([]Feed, error) {
	query := "SELECT id, latest_post_guid, discord_channel_id, feed_url FROM feeds"

	rows, err := DB.Query(query)
	if err != nil {
		log.Error("Error getting feeds.")
		return nil, err
	}
	defer rows.Close()

	var feeds []Feed
	for rows.Next() {
		var feed Feed
		err := rows.Scan(&feed.ID, &feed.LatestPostGUID, &feed.DiscordChannelID, &feed.FeedURL)
		if err != nil {
			log.Error("Error scanning row.", err)
			continue
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
