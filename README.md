# tg-shortener
Telegram bot for url shortener project (https://github.com/OmelchenkoVladimir/url-shortener)

![alt text](https://github.com/OmelchenkoVladimir/tg-shortener/blob/main/examples/static/TgBotScreen.png?raw=true)

Launch instructions:
1. Make sure url-shortener is up and running at HOST:PORT (you can change values in configuration file)
2. Generate telegram token for your telegram bot
3. Set OS environment variable TG_TOKEN to this token
4. Use command [go run main.go] (without brackets) in your terminal

Now you should be able to connect to your telegram bot. Send /encode <link> (without brackets) to get short link.
