package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, "You've typed start command")

	_, err := b.bot.Send(outMsg)
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
	outMsg.ReplyToMessageID = inMsg.MessageID

	_, err := b.bot.Send(outMsg)
	return err
}
