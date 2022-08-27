package commands

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const LIMIT = 20

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	list, done := c.productService.List(0, LIMIT)
	outputMsg := "Here are all the products:\n\n"
	for _, p := range list {
		outputMsg += fmt.Sprintf("%d. %s\n", p.ID, p.Title)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	newOffset := LIMIT
	if done {
		newOffset = 0
	}
	bytes, _ := json.Marshal(CommandData{Offset: newOffset})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(bytes)),
		))

	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}

func (c *Commander) NextPage(cq *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}
	if err := json.Unmarshal([]byte(cq.Data), &parsedData); err != nil {
		msg := tgbotapi.NewMessage(
			cq.Message.Chat.ID,
			"Произошла внутренняя ошибка, попробуйте снова")
		if _, err := c.bot.Send(msg); err != nil {
			log.Println("error when send message to tg", err)
		}
	}

	list, done := c.productService.List(parsedData.Offset, LIMIT)
	outputMsg := "Here are all the products:\n\n"
	for _, p := range list {
		outputMsg += fmt.Sprintf("%d. %s\n", p.ID, p.Title)
	}

	msg := tgbotapi.NewMessage(cq.Message.Chat.ID, outputMsg)

	newOffset := parsedData.Offset + LIMIT
	if done {
		newOffset = 0
	}
	bytes, _ := json.Marshal(CommandData{Offset: newOffset})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(bytes)),
		))

	if _, err := c.bot.Send(msg); err != nil {
		log.Println("error when send message to tg", err)
	}
}
