package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siteddv/pocketel_bot/pkg/repository"
)

func (b *Bot) initAuthorizationProcess(inMsg *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(inMsg.Chat.ID)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf(replyStartTemplate, authLink)

	outMsg := tgbotapi.NewMessage(inMsg.Chat.ID, msgText)

	_, err = b.bot.Send(outMsg)
	return err
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	result, err := b.tokenRepository.Get(chatID, repository.AccessTokens)
	return result, err
}

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectUrl(chatID)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrl)
	if err != nil {
		return "", err
	}

	err = b.tokenRepository.Save(chatID, requestToken, repository.RequestTokens)
	if err != nil {
		return "", err
	}

	authLink, err := b.pocketClient.GetAuthorizationURL(requestToken, redirectUrl)

	return authLink, err
}

func (b *Bot) generateRedirectUrl(chatID int64) string {
	result := fmt.Sprintf("%s?chat_id=%d", b.redirectUrl, chatID)
	return result
}
