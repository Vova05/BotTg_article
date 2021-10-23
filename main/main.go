package main

import (
	tgbotapi "Projects/awesomeProject/main/telegram-bot-api-master"
	"bytes"
	//tgbotapi "telegram-bot-api-master"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	array_of_commands:=DB_command()

	for update := range updates {

		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text,update.Message.Chat.ID)

		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply := "Привет. Я телеграм-бот"
			// создаем ответное сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			// отправляем
			bot.Send(msg)
		case "hello":
			reply := "world"
			// создаем ответное сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			// отправляем
			bot.Send(msg)
		case "info":
			// отправляем
			buffer := bytes.Buffer{}
			tmp :=1
			for _,val:=range array_of_commands{
				buffer.WriteString(val)
				if tmp==1{
					buffer.WriteString(" - ")
					tmp++
				}else{
					buffer.WriteString("\n")
					tmp--
				}
			}
			msg :=tgbotapi.NewMessage(update.Message.Chat.ID,buffer.String())
			bot.Send(msg)
		case "create_article":
			reply := "Новая статья успешно добавлена"
			if Create_article(update.Message.Text)==3{
				reply="Не корректное имя"
			}
			msg :=tgbotapi.NewMessage(update.Message.Chat.ID,reply)
			bot.Send(msg)
		case "save_link":
			reply := "Новая ссылка успешно добавлена"
			if Create_link_article(update.Message.Text)==3{
				reply="Ссылка уже есть"
			}
			msg :=tgbotapi.NewMessage(update.Message.Chat.ID,reply)
			bot.Send(msg)
		}

		switch update.Message.Chat.ID {
		case -758942813:
			switch update.Message.Command() {
			case "sos":
				msg := tgbotapi.NewMessage(940757224, update.Message.Text + " \n Из чата " + update.Message.Chat.Title + "\n От "+ update.Message.From.UserName )
				bot.Send(msg)
			}
		}

	}
}




