package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	list := c.productService.List()
	outputMsg := "Here are all the products:\n\n"
	for _, p := range list {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
