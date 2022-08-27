package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Panic(inputMessage *tgbotapi.Message) {
	panic("AAAAAAAA!!!!")
}
