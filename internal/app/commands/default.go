package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Вы написали: "+inputMessage.Text)
	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
