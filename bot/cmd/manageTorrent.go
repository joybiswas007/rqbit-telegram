package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joybiswas007/rqbit-telegram/bot/botutils"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

// add magnet link or any http torrent link
func ManageTorrent(b *gotgbot.Bot, ctx *ext.Context) error {
	conf := utils.GetConfig()
	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	apiPath := map[string]string{
		"/pause":  "pause",  //pause a torrent
		"/resume": "start",  // resume a torrent
		"/delete": "delete", // delete a torrent from client with files
		"/remove": "forget", // just remove the torrent from client
	}

	msg := map[string]string{
		"pause":  "paused",
		"resume": "started",
		"delete": "deleted",
		"forget": "removed",
	}

	parts := strings.Split(ctx.Message.Text, " ")
	if len(parts) != 2 {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>id can't be empty. Must provide valid id to %s the job.</b>", apiPath[parts[0]]))
		return nil
	}

	tID, err := strconv.Atoi(parts[1])
	if err != nil {
		botutils.ReplyMessage(b, ctx, "<b>invalid id :(</b>")
		return nil
	}

	torrent, err := conf.RQBit.GetTorrentInfo(tID)
	if err != nil {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", err))
		return nil
	}

	if err := conf.RQBit.ManageTorrent(tID, apiPath[parts[0]]); err != nil {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", err))
		return nil
	}
	botutils.ReplyMessage(b, ctx, fmt.Sprintf("<code>%s</code> %s successfully.", torrent.Name, msg[apiPath[parts[0]]]))

	return nil
}
