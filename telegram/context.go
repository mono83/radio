package telegram

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mono83/radio"
)

const SCHEMA = "telegram"

type context struct {
	bot     *tgbotapi.BotAPI
	message tgbotapi.Message
}

func (c context) GetUser() radio.User {
	return radio.User{
		Schema: SCHEMA,
		ID:     strconv.FormatInt(c.message.Chat.ID, 10),
	}
}

func (c context) GetArgs() []string {
	return strings.Split(strings.TrimSpace(c.message.Text), " ")
}

func (c context) CommandInProgress() {
	_, _ = c.bot.Send(tgbotapi.NewChatAction(c.message.Chat.ID, tgbotapi.ChatTyping))
}

func (c context) SendMessage(msg interface{}) {
	if s, ok := msg.(string); ok {
		c.stringMessage(s)
	}
	if e, ok := msg.(error); ok {
		c.errorMessage(e)
	}
}

func (c context) stringMessage(str string) {
	_, _ = c.bot.Send(tgbotapi.NewMessage(c.message.Chat.ID, str))
}

func (c context) errorMessage(e error) {
	msg := tgbotapi.NewMessage(c.message.Chat.ID, "`"+e.Error()+"`")
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, _ = c.bot.Send(msg)
}
