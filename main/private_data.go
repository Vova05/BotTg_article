package main

import (
	"flag"
	"log"
	"os"
)

var (
	telegramBotToken string
)

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "1901931911:AAFINPGIJ4psmAvUJFrpTm4dqaoj-zcog9Q", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}
