package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/siteddv/pocketel_bot/pkg/telegram"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Put your bot token by "botToken" key into file ".env"
	botToken := os.Getenv("botToken")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	if err = telegramBot.Start(); err != nil {
		log.Fatalf("Error during starting bot")
	}
}
