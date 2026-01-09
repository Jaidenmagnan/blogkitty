package commands

import "github.com/bwmarrin/discordgo"

func Ping(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "meow",
		},
	})

}
