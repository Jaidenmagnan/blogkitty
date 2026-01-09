# this scrip will run the bot in the background
go build -o blogkitty ./
nohup ./blogkitty >blogkitty.log 2>&1 &

# to view the logs do this to the view
# tail -f bot.log
