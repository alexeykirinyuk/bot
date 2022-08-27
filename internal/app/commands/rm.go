package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Rm(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(arg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Не получилось распознать индекс")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
		return
	}

	if err := c.productService.Rm(idx); err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Неправильный индекс: "+err.Error())
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
		return
	}

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт '"+arg+"' успешно удален")
	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
