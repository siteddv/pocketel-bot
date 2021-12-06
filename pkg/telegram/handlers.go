package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	startCommand = "start"

	replyStartTemplate = "Hello. To save links in your Pocket account you need to provide access me to it. Please follow this link:\n%s"
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
	authLink, err := b.generateAuthorizationLink(inMsg.Chat.ID)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf(replyStartTemplate, authLink)

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, msgText)

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
	outMsg.ReplyToMessageID = inMsg.MessageID

	_, err := b.bot.Send(outMsg)
	return err
}
