package main

import (
	"github.com/Jaidenmagnan/blogkitty/commands"
	"github.com/Jaidenmagnan/blogkitty/db"
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	godotenv.Load()

	err := db.Connect()
	if err != nil {
		log.Fatal("could not connect to the db")
	}

	token := os.Getenv("DISCORD_TOKEN")
	guildID := os.Getenv("GUILD_ID")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Error("error initializing the discord bot")
		return
	}

	// We can put all of our event handlers here.
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// We must create each of our commands here anad add the handlers.
	commandList := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "responds with pong",
		},
		{
			Name:        "monitor",
			Description: "Add a blog feed to your server.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "blog-feed",
					Description: "The link to the rss feed of the blog you would like to monitor.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "channel-name",
					Description: "The name of the channel to be created which blog posts will be sent in.",
					Required:    true,
				},
			},
		},
	}

	commandHandlers := map[string]func(dg *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":    commands.Ping,
		"monitor": commands.Monitor,
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	err = dg.Open()
	if err != nil {
		log.Error("could not open ws connection")
		return
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commandList))
	for i, v := range commandList {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, v)
		if err != nil {
			log.Error("Cannot create '%v' command: %v", "name", v.Name, "error", err)
			return
		}
		log.Info("Registered the command: ", "name", v.Name)
		registeredCommands[i] = cmd
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
