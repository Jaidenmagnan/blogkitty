package commands

import (
	"github.com/Jaidenmagnan/blogkitty/db"
	"github.com/Jaidenmagnan/blogkitty/rss"
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

func reply(message string, dg *discordgo.Session, i *discordgo.InteractionCreate) {
	dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func Monitor(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	rssUrl := i.ApplicationCommandData().GetOption("rss-url").StringValue()
	channelName := i.ApplicationCommandData().GetOption("channel-name").StringValue()

	log.Info("blog feed", "blog", rssUrl)
	log.Info("channel name", "channel", channelName)

	ok := rss.IsFeed(rssUrl)
	if !ok {
		msg := "Could not parse feed."
		log.Error(msg)
		reply(msg, dg, i)
		return
	}

	// TODO: Add the new channel under a particular category.
	channel, err := dg.GuildChannelCreate(i.GuildID, channelName, discordgo.ChannelTypeGuildText)
	if err != nil {
		msg := "Could not create channel."
		log.Error(msg)
		reply(msg, dg, i)
		return
	}

	feed := db.Feed{
		DiscordChannelID: channel.ID,
		FeedURL:          rssUrl,
		LatestPostGUID:   "",
	}

	feed, err = db.InsertFeed(feed)
	if err != nil {
		msg := "Could not add the feed to the DB."
		log.Error(msg)
		reply(msg, dg, i)
	}

	reply("Successfully created the channel **"+channel.Name+"** to monitor your feed!", dg, i)

	dg.ChannelMessageSend(channel.ID, "Meow! This channel be monitoring *"+feed.FeedURL+"*.")
}
