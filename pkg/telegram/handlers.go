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

	replyStartTemplate     = "Hello. To save links in your Pocket account you need to provide access me to it. Please follow this link:\n%s"
	replyAlreadyAuthorized = "You've already authorized. Please send me a link and I'll save it to Pocket"
	replySuccessfulSave    = "The link successfully saved"
	replyInvalidLink       = "This is an invalid link"
	replyUnauthorized      = "You aren't authorized. Please use the command /start"
	replyInternalError     = "I couldn't save link. Please try again"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
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

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(outMsg)

	return err
}

func (b *Bot) handleUnknownCommand(inMsg *tgbotapi.Message) error {
	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, "You've typed the command that I don't know")

	_, err := b.bot.Send(outMsg)
	return err
}

func (b *Bot) handleMessage(inMsg *tgbotapi.Message) error {
	log.Printf("[%s] %s", inMsg.From.UserName, inMsg.Text)

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, inMsg.Text)

	if _, err := url.ParseRequestURI(inMsg.Text); err != nil {
		outMsg.Text = replyInvalidLink
		_, err = b.bot.Send(outMsg)
		return err
	}

	accesstoken, err := b.getAccessToken(inMsg.Chat.ID)
	if err != nil {
		outMsg.Text = replyUnauthorized
		_, err = b.bot.Send(outMsg)
		return err
	}
	addInput := pocket.AddInput{
		AccessToken: accesstoken,
		URL:         inMsg.Text,
	}
	if err = b.pocketClient.Add(context.Background(), addInput); err != nil {
		outMsg.Text = replyInternalError
		_, err = b.bot.Send(outMsg)
	}

	_, err = b.bot.Send(outMsg)

	return err
}
