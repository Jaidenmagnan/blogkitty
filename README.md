# blog-kitty
Hi! Welcome to blog kitty!

This discord bot written in [discordgo](https://github.com/bwmarrin/discordgo) that will simply monitor rss feeds in your server.

You can add the bot to your server (link)[will be added soon].

# Usage
To use this bot, first either add it to your server or run it locally. Information for running the bot locally can be found below.

Simply just use the following slash command in Discord.

```
/monitor [rss-feed] [channel-name]
```

- The `rss-feed` parameter is a link to the rss feed you would wish to monitor.
- The `channel-name` parameter is the name of the channel you would like to create that will monitor the feed.

To remove the feed, just delete the channel. If you want to add the command yourself go ahead!! Always welcoming contributions.

# Local Usage
To run the bot locally you must get your own `DISCORD_TOKEN` from the Discord developers portal. I am not going to go into too much detail into getting the bot setup on the Discord side. I recommend [this](https://discordjs.guide/legacy/preparations/adding-your-app) documentation from discord.js (yes, even though this is Go).

Then to run the migrations locally I use [Goose](https://github.com/pressly/goose). So just run the following:

```
go get -tool https://github.com/pressly/goose
go tool goose up 
```

This will run the migrations. Then just run the bot.
```
go run .
```

Or you can compile it and run it.
```
go build
./blogkitty
```

