package telegram

import (
	"context"
	"fmt"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectUrl(chatID)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrl)
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
