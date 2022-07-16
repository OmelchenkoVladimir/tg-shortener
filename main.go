package main

import (
	"fmt"
	"log"
	"os"
	"tg_shortener/config"
	"tg_shortener/conn"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	TOKEN := os.Getenv("TG_TOKEN")
	conf, err := config.ReadConfig("./config/config.json")
	if err != nil {
		log.Fatal("Can't open config")
	}
	// kinda dodn't want to use url [name collision]
	urlShortenerPath := "http://" + conf.Url_shortener.Host + ":" + conf.Url_shortener.Port + "/encode/"

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic("Error: can't connect to my bot;", err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				chatId := update.Message.Chat.ID
				fmt.Println("GOT COMMAND!")
				switch update.Message.Command() {
				case "hello":
					reply := tgbotapi.NewMessage(chatId, "Hi!")
					bot.Send(reply)
				case "shorturl":
					var reply tgbotapi.MessageConfig
					url := update.Message.CommandArguments()
					result, err := conn.GetShortUrl(url, urlShortenerPath)
					if err != nil {
						log.Println(err.Error())
						reply = tgbotapi.NewMessage(chatId, "Sorry, failed to do it :(")
					} else {
						reply = tgbotapi.NewMessage(chatId, result)
					}
					bot.Send(reply)
				}
			} else {
				log.Printf("Message from %s! %s", update.Message.From.UserName, update.Message.Text)
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Thanks for telling me that")
				reply.ReplyToMessageID = update.Message.MessageID
				bot.Send(reply)
			}
		}
	}
}
