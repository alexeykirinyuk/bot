package commands

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Update(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	args := strings.Split(arg, " ")
	if len(args) != 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Кол-во слов должно быть равно двум")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
		return
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Не получилось распознать индекс")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
		return
	}

	if err := c.productService.Update(idx, args[1]); err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Неправильный индекс: "+err.Error())
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
		return
	}

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт '"+args[0]+"' успешно обновлен")
	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
