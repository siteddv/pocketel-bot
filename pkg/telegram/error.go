package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siteddv/pocketel_bot/pkg/config"
)

var (
	errInvalidUrl   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save link")
)

func (b *Bot) handleError(chatID int64, err error, cfg *config.Config) {
	msg := tgbotapi.NewMessage(chatID, cfg.Messages.Errors.Default)

	switch err {
	case errInvalidUrl:
		msg.Text = cfg.Messages.Errors.InvalidLink
	case errUnauthorized:
		msg.Text = cfg.Messages.Errors.Unauthorized
	case errUnableToSave:
		msg.Text = cfg.Messages.Errors.UnableToSave
	}

	b.bot.Send(msg)
}
