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
	query := "INSERT INTO feeds (feed_url, discord_channel_id) VALUES (?, ?)"

	result, err := DB.Exec(query, feed.FeedURL, feed.DiscordChannelID, feed.LatestPostGUID)
	if err != nil {
		log.Error("error inserting feed", err)
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
