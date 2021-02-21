package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kalyasik/gmail_telegram_bot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
