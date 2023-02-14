package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	for _, member := range Members {
		if now.Format("01-02") == member.Birthday {
			if _, err := bot.Send(tgbotapi.NewMessage(chatId, member.Name+"님의 생일을 축하합니다🎉")); err != nil {
				panic(err)
			}
		}
	}
}
