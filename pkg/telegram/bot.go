package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siteddv/golang-pocket-sdk"
	"log"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client) *Bot {
	return &Bot{
		bot:          bot,
		pocketClient: pocketClient,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)

	return updates, err
}
