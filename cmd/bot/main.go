package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/siteddv/golang-pocket-sdk"
	"github.com/siteddv/pocketel_bot/pkg/telegram"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error handled during loading env variables: %s", err.Error())
	}

	// Put your bot token by "botToken" key into file ".env"
	botToken := os.Getenv("botToken")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("error handled during creating telegram bot client: %s", err.Error())
	}

	bot.Debug = true

	// Put your consumer key by "consumerKey" key into file ".env"
	consumerKey := os.Getenv("consumerKey")
	client, err := pocket.NewClient(consumerKey)
	if err != nil {
		log.Fatalf("error handled during creating pocket client: %s", err.Error())
	}

	_, err = bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatalf("error handled during initializing db: %s", err.Error())
	}

	telegramBot := telegram.NewBot(bot, client, "google.com")
	if err = telegramBot.Start(); err != nil {
		log.Fatalf("Error during starting bot")
	}
}
