package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Panic("token not found - TELEGRAM_TOKEN")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 60,
	})

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if strings.EqualFold(update.Message.Text, "нет") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну пожалуйста...")
			bot.Send(msg)
		} else if strings.EqualFold(update.Message.Text, "да") || strings.EqualFold(update.Message.Text, "ладно") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ура")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Почитаешь мне книжку вечером?")
			bot.Send(msg)
		}

	}
}
