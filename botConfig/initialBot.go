package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	// Инициализация бота с использованием токена
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}
	_, err = bot.RemoveWebhook()
	if err != nil {
		log.Panic(err)
	}

	// Установка режима длинных запросов (long polling)
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	// Проход по каналу обновлений
	for update := range updates {
		if update.Message == nil { // игнорировать обновления, не связанные с сообщениями
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Ответ на сообщение
		reply := "Привет, я бот! Ты написал: " + update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
