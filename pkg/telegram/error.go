package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var (
	errInvalidUrl   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save link")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Errors.Default)

	switch err {
	case errInvalidUrl:
		msg.Text = b.messages.Errors.InvalidLink
	case errUnauthorized:
		msg.Text = b.messages.Errors.Unauthorized
	case errUnableToSave:
		msg.Text = b.messages.Errors.UnableToSave
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		log.Printf("Unable to send message due to error: %s", err.Error())
	}
}
