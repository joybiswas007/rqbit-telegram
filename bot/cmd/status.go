package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joybiswas007/rqbit-telegram/bot/botutils"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

func Status(b *gotgbot.Bot, ctx *ext.Context) error {
	conf := utils.GetConfig()

	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	parts := strings.Split(ctx.Message.Text, " ")
	if len(parts) != 2 {
		botutils.ReplyMessage(b, ctx, "<b>id can't be empty. Must provide valid id</b>")
		return nil
	}

	tId, _ := strconv.Atoi(parts[1])

	stats, err := conf.RQBit.GetTorrentStats(tId)
	if err != nil {
		log.Printf("%v", err)
	}

	message := botutils.FormatTorrentStats(stats)
	botutils.ReplyMessage(b, ctx, message)
	return nil
}
