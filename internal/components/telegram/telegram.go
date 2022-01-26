package main

import (
	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"infrastructure-telegram/internal/components/utils"
	"time"
)

type Telegram interface {
	Start() error
	Send(msg string)
}

type telegram struct {
	online   bool
	messages chan message
	// Клиент Telegram BotAPI
	bot *tgbotapi.BotAPI
}

func (tg telegram) Send(text string) {
	tg.messages <- message{
		text: text,
	}
}

type message struct {
	chatID int64
	text   string
}

func New() Telegram {
	telegram := &telegram{
		online:   false,
		messages: make(chan message, 50),
	}

	go telegram.daemonSender()

	return telegram
}

func (tg telegram) daemonSender() {
	timer := time.NewTicker(time.Second / 30)
	for range timer.C {
		if tg.online {
			message := <-tg.messages
			go func() {
				tgmsg := tgbotapi.NewMessage(message.chatID, message.text)
				tgmsg.ParseMode = "TEXT"
				// Собственно, отправка сообщения.
				if _, err := tg.bot.Send(tgmsg); err != nil {
					utils.CaptureEvent(sentry.Event{
						Message: "Ошибка отправки сообщения в Telegram",
						Extra: map[string]interface{}{
							"ChatID": tgmsg.ChatID,
							"Text":   tgmsg.Text,
						},
					})
				}
			}()
		}
	}
}

func (tg telegram) Start() error {
	panic("implement me")
}
