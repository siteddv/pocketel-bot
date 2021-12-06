package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/siteddv/golang-pocket-sdk"
	"github.com/siteddv/pocketel_bot/pkg/repository"
	"github.com/siteddv/pocketel_bot/pkg/repository/boltdb"
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

	db, err := initDB()
	if err != nil {
		log.Fatalf("error handled during initing database: %s", err.Error())
	}

	tokenRepos := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, client, tokenRepos, "google.com")
	if err = telegramBot.Start(); err != nil {
		log.Fatalf("Error during starting bot: %s", err.Error())
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(
		func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
			if err != nil {
				return err
			}

			_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
			if err != nil {
				return err
			}

			return nil
		},
	)

	return db, err
}
