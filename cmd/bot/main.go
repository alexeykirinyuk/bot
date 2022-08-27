package main

import (
	"log"
	"os"

	"github.com/alexeykirinyuk/bot/internal/app/commands"
	"github.com/alexeykirinyuk/bot/internal/service/product"
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

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 60,
	})

	for update := range updates {
		commander.HandleUpdate(&update)
	}
}
