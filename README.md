# tg-shortener
Telegram bot for url shortener project (https://github.com/OmelchenkoVladimir/url-shortener)

Launch instructions:
1. Make sure url-shortener is up and running as HOST:PORT (you can change values in configuration file)
2. Generate telegram token for your telegram bot
3. Set OS enviroment variable TG_TOKEN to this token
4. go run main.go

Now you should be able to connect to your telegram bot. Send /encode <link> (without brackets) to get short link.
