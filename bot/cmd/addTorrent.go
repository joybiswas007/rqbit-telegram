package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joybiswas007/rqbit-telegram/bot/botutils"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

// add magnet link or any http torrent link
func AddTorrent(b *gotgbot.Bot, ctx *ext.Context) error {
	conf := utils.GetConfig()

	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	parts := strings.Split(ctx.Message.Text, " ")
	if len(parts) != 2 {
		botutils.ReplyMessage(b, ctx, "<b>Command can't be empty. Must provide valid magnet link or http torrent link.</b>")
		return nil
	}

	meta, err := conf.RQBit.AddTorrent(parts[1])
	if err != nil {
		log.Printf("%v", err)
	}

	botutils.ReplyMessage(b, ctx, fmt.Sprintf("Torrent: <code>%s</code> have been added to the client successfully. Check status using <code>/status</code>", meta.Details.Name))

	// delete the magnet link or http torrent link from chat after it's been added to the client
	ok, err := b.DeleteMessage(ctx.EffectiveChat.Id, ctx.EffectiveMessage.MessageId, &gotgbot.DeleteMessageOpts{})
	if !ok || err != nil {
		log.Printf("%v", err)
	}

	return nil
}
