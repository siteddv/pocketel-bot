package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const ()

var (
	errInvalidUrl   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save link")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, replyUnknownError)

	switch err {
	case errInvalidUrl:
		msg.Text = replyInvalidLink
	case errUnauthorized:
		msg.Text = replyUnauthorized
	case errUnableToSave:
		msg.Text = replyInternalError
	}

	b.bot.Send(msg)
}
