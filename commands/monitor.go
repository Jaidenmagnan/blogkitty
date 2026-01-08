package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

func Monitor(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	blogFeed := i.ApplicationCommandData().GetOption("blog-feed").StringValue()
	channelName := i.ApplicationCommandData().GetOption("channel-name").StringValue()

	log.Info("blog feed", "blog", blogFeed)
	log.Info("channel name", "channel", channelName)

	dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "check logs",
		},
	})
}
