package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(arg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Не получилось распознать индекс")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
	}

	prod, err := c.productService.Get(idx)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "product not found")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "product: "+prod.Title)
	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
