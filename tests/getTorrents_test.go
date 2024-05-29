package tests

import (
	"os"
	"testing"

	"github.com/joybiswas007/rqbit-telegram/utils"
)

func TestGetTorrents(t *testing.T) {
	utils.LoadConfig("../config.toml")
	conf := utils.GetConfig()

	torrents, err := conf.RQBit.GetTorrents()
	if err != nil {
		t.Logf("%v", err)
	}
	if len(torrents.Torrents) == 0 {
		t.Log("queue is empty")
		os.Exit(1)
	}

	for _, torrent := range torrents.Torrents {
		t.Log(torrent.TorrentID)
		t.Log(torrent.InfoHash)
	}
}
