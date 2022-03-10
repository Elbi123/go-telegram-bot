package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var starterKey = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("7"),
		tgbotapi.NewKeyboardButton("8"),
		tgbotapi.NewKeyboardButton("9"),
		tgbotapi.NewKeyboardButton("*"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
		tgbotapi.NewKeyboardButton("/"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
		tgbotapi.NewKeyboardButton("-"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("0"),
		tgbotapi.NewKeyboardButton("."),
		tgbotapi.NewKeyboardButton("="),
		tgbotapi.NewKeyboardButton("+"),
	),
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	telegram_apitoken := os.Getenv("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(telegram_apitoken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.Text = update.Message.Time().Format("dd-mm-yyyy")
			msg.ReplyMarkup = tgbotapi.NewInlineQueryResultArticle("123", "Hello", "This is new inline query result article")

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.InlineQuery.Query == "" {

			fmt.Println("===========================")
			fmt.Println(update.InlineQuery.Query)
			fmt.Println("===========================")
			article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", update.Message.Text)

			inlineConfig := tgbotapi.InlineConfig{
				InlineQueryID: update.InlineQuery.ID,
				IsPersonal:    true,
				CacheTime:     0,
				Results:       []interface{}{article},
			}

			if _, err := bot.Send(inlineConfig); err != nil {
				panic(err)
			}
		}

	}
}
