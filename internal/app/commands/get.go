package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(arg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "invalid args")
		c.bot.Send(msg)
		return
	}

	prod, err := c.productService.Get(idx)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "product not found")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "product: "+prod.Title)
	c.bot.Send(msg)
}
