package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get - get products. example - '/get 1'\n"+
			"/add - add product. example - '/add my-dear-product'\n"+
			"/update - update product. example - '/update 2 my-dear-product-2'\n"+
			"/rm - remove product. example - 'rm 1'")
	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
