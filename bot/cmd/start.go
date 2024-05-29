package cmd

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	buttons := [][]gotgbot.InlineKeyboardButton{
		{
			{
				Text: "Developer",
				Url:  "https://github.com/joybiswas007",
			},
			{
				Text: "Source code",
				Url:  "https://github.com/joybiswas007/rqbit-telegram",
			},
		},
	}

	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s. A Telegram bot to remotely control rqbit instance.", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode:   "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons},
	})

	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	return nil
}
