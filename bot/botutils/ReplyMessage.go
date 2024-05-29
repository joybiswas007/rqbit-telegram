package botutils

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// Utility function to reply to user messages
func ReplyMessage(b *gotgbot.Bot, ctx *ext.Context, message string) (*gotgbot.Message, error) {
	msg, err := ctx.EffectiveMessage.Reply(b, message, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return &gotgbot.Message{}, err
	}
	return msg, nil
}
