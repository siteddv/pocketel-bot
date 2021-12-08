package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pocket "github.com/siteddv/golang-pocket-sdk"
	"log"
	"net/url"
)

const (
	startCommand = "start"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) handleCommand(inMsg *tgbotapi.Message) error {
	switch inMsg.Command() {
	case startCommand:
		return b.handleStartCommand(inMsg)
	default:
		return b.handleUnknownCommand(inMsg)
	}
}

func (b *Bot) handleStartCommand(inMsg *tgbotapi.Message) error {
	_, err := b.getAccessToken(inMsg.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(inMsg)
	}

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, b.messages.Responses.AlreadyAuthorized)
	_, err = b.bot.Send(outMsg)

	return err
}

func (b *Bot) handleUnknownCommand(inMsg *tgbotapi.Message) error {
	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, b.messages.Responses.UnknownCommand)

	_, err := b.bot.Send(outMsg)
	return err
}

func (b *Bot) handleMessage(inMsg *tgbotapi.Message) error {
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

	outMsg.Text = b.messages.Responses.SavedSuccessfully
	_, err = b.bot.Send(outMsg)

	return err
}
