package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pocket "github.com/siteddv/golang-pocket-sdk"
	"github.com/siteddv/pocketel_bot/pkg/config"
	"github.com/siteddv/pocketel_bot/pkg/repository"
	"github.com/siteddv/pocketel_bot/pkg/repository/boltdb"
	"github.com/siteddv/pocketel_bot/pkg/server"
	"github.com/siteddv/pocketel_bot/pkg/telegram"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		return
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("error handled during creating telegram bot client: %s", err.Error())
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.ConsumerKey)
	if err != nil {
		log.Fatalf("error handled during creating pocket client: %s", err.Error())
	}

	db, err := initDB(cfg.DbPath)
	if err != nil {
		log.Fatalf("error handled during initing database: %s", err.Error())
	}

	tokenRepos := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepos, cfg.AuthServerUrl, cfg.Messages)

	authServer := server.NewAuthorizationServer(pocketClient, tokenRepos, cfg.BotUrl)

	go func() {
		if err = telegramBot.Start(); err != nil {
			log.Fatalf("error handled during starting bot: %s", err.Error())
		}
	}()

	if err := authServer.Start(); err != nil {
		log.Fatalf("error handled during server starting: %s", err.Error())
	}
}

func initDB(dbPath string) (*bolt.DB, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
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
