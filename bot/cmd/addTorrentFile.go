package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joybiswas007/rqbit-telegram/bot/botutils"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

// send any valid .torrent file to the chat bot will add it to client
func AddTorrentFile(b *gotgbot.Bot, ctx *ext.Context) error {
	conf := utils.GetConfig()

	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	doc := ctx.EffectiveMessage.Document
	if doc.MimeType != "application/x-bittorrent" {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("File: <code>%s</code> is invalid. Try again with a valid <code>.torrent</code> file.", doc.FileName))
		return nil
	}

	f, err := b.GetFile(doc.FileId, &gotgbot.GetFileOpts{})
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	downloadsDir := filepath.Join(pwd, "torrents")
	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		os.Mkdir(downloadsDir, 0755)
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", conf.Bot.Token, f.FilePath)

	filePath := filepath.Join(downloadsDir, doc.FileName)

	torrentPath, err := botutils.SaveTorrent(fileURL, filePath)
	if err != nil {
		log.Printf("%v", err)
	}

	meta, err := conf.RQBit.AddTorrent(torrentPath)
	if err != nil {
		log.Printf("%v", err)
	}

	botutils.ReplyMessage(b, ctx, fmt.Sprintf("Torrent: <code>%s</code> have been added to the client successfully. Check status using <code>/status</code>", meta.Details.Name))

	//delete the torrent file from locally
	if err := os.RemoveAll(torrentPath); err != nil {
		log.Printf("failed to delete torrent: %v", err)
	}

	// delete the .torrent file from chat after it's been added to the client
	ok, err := b.DeleteMessage(ctx.EffectiveChat.Id, ctx.EffectiveMessage.MessageId, &gotgbot.DeleteMessageOpts{})
	if !ok || err != nil {
		log.Printf("%v", err)
	}

	return nil
}
