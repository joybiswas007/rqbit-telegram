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

// get all task id to and use them to view specific task details
func GetTorrentIDs(b *gotgbot.Bot, ctx *ext.Context) error {
	conf := utils.GetConfig()
	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	torrents, err := conf.RQBit.GetTorrents()
	if err != nil {
		log.Printf("%v", err)
	}

	if len(torrents.Torrents) == 0 {
		botutils.ReplyMessage(b, ctx, "<b>No torrent ID(s) found  :(</b>")
		return nil
	}

	var builder strings.Builder

	builder.WriteString("<b>Torrent IDs:</b>\n\n")
	for _, torrent := range torrents.Torrents {
		builder.WriteString("<code>")
		builder.WriteString(fmt.Sprintf("/status %d", torrent.TorrentID))
		builder.WriteString("</code>\n")
	}
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("<b>Total torrents: %d</b>", len(torrents.Torrents)))

	message := builder.String()

	botutils.ReplyMessage(b, ctx, message)

	return nil
}
