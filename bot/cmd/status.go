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

	torrent, err := conf.RQBit.GetTorrentInfo(tId)
	if err != nil {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", err))
		return nil
	}

	stats, err := conf.RQBit.GetTorrentStats(tId)
	if err != nil {
		log.Printf("%v", err)
	}
	timeRemaining := "NaN"
	if stats.Live.TimeRemaining.HumanReadable != "" {
		timeRemaining = fmt.Sprintf("%v", stats.Live.TimeRemaining.HumanReadable)
	}

	message := fmt.Sprintf(
		"<b>Torrent Status:</b>\n"+
			"Name: <code>%s</code>\n"+
			"InfoHash: <code>%s</code> \n"+
			"State: %s\n"+
			"Progress: %s / %s\n"+
			"Uploaded: %s \n"+
			"Finished: %t\n"+
			"Download Speed: %s\n"+
			"Upload Speed: %s\n"+
			"Peers: %d seen, %d live\n"+
			"Average Piece Download Time: %ds %dns\n"+
			"Time Remaining: %s",
		torrent.Name,
		torrent.InfoHash,
		stats.State,
		utils.HumanReadableBytes(uint64(stats.ProgressBytes)),
		utils.HumanReadableBytes(uint64(stats.TotalBytes)),
		utils.HumanReadableBytes(uint64(stats.UploadedBytes)),
		stats.Finished,
		stats.Live.DownloadSpeed.HumanReadable,
		stats.Live.UploadSpeed.HumanReadable,
		stats.Live.Snapshot.PeerStats.Seen,
		stats.Live.Snapshot.PeerStats.Live,
		stats.Live.AveragePieceDownloadTime.Secs,
		stats.Live.AveragePieceDownloadTime.Nanos,
		timeRemaining,
	)
	botutils.ReplyMessage(b, ctx, message)
	return nil
}
