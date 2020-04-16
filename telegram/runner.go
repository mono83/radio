package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mono83/radio"
)

// Run starts Telegram over given data connector
func Run(token string, c radio.Connector) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Building context
		ctx := &context{
			bot:     bot,
			message: *update.Message,
		}

		// Running
		if err := c.Invoke(ctx); err != nil {
			log.Print(err)
		}
	}
	return nil
}
