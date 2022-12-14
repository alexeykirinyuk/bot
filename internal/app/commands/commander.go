package commands

import (
	"fmt"
	"log"

	"github.com/alexeykirinyuk/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, inputMessage *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Println("recovered from panic:", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		c.NextPage(update.CallbackQuery)
		return
	}

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	case "add":
		c.Add(update.Message)
	case "update":
		c.Update(update.Message)
	case "rm":
		c.Rm(update.Message)
	default:
		c.Default(update.Message)
	}
}
