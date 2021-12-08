package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pocket "github.com/siteddv/golang-pocket-sdk"
	"github.com/siteddv/pocketel_bot/pkg/config"
	"log"
	"net/url"
)

const (
	startCommand = "start"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, cfg *config.Config) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message, cfg); err != nil {
				b.handleError(update.Message.Chat.ID, err, cfg)
			}
			continue
		}

		if err := b.handleMessage(update.Message, cfg); err != nil {
			b.handleError(update.Message.Chat.ID, err, nil)
		}
	}
}

func (b *Bot) handleCommand(inMsg *tgbotapi.Message, cfg *config.Config) error {
	switch inMsg.Command() {
	case startCommand:
		return b.handleStartCommand(inMsg, cfg)
	default:
		return b.handleUnknownCommand(inMsg, cfg)
	}
}

func (b *Bot) handleStartCommand(inMsg *tgbotapi.Message, cfg *config.Config) error {
	_, err := b.getAccessToken(inMsg.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(inMsg, cfg)
	}

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, cfg.Messages.Responses.AlreadyAuthorized)
	_, err = b.bot.Send(outMsg)

	return err
}

func (b *Bot) handleUnknownCommand(inMsg *tgbotapi.Message, cfg *config.Config) error {
	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, cfg.Messages.Responses.UnknownCommand)

	_, err := b.bot.Send(outMsg)
	return err
}

func (b *Bot) handleMessage(inMsg *tgbotapi.Message, cfg *config.Config) error {
	log.Printf("[%s] %s", inMsg.From.UserName, inMsg.Text)

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, inMsg.Text)

	if _, err := url.ParseRequestURI(inMsg.Text); err != nil {
		return errInvalidUrl
	}

	accessToken, err := b.getAccessToken(inMsg.Chat.ID)
	if err != nil {
		return errUnauthorized
	}
	addInput := pocket.AddInput{
		AccessToken: accessToken,
		URL:         inMsg.Text,
	}
	if err = b.pocketClient.Add(context.Background(), addInput); err != nil {
		return errUnableToSave
	}

	outMsg.Text = cfg.Messages.Responses.SavedSuccessfully
	_, err = b.bot.Send(outMsg)

	return err
}
