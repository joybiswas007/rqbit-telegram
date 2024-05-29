package tests

import (
	"testing"

	"github.com/joybiswas007/rqbit-telegram/utils"
)

func TestGetTorrentStats(t *testing.T) {
	utils.LoadConfig("../config.toml")
	conf := utils.GetConfig()

	stats, err := conf.RQBit.GetTorrentStats(3)
	if err != nil {
		t.Log(err)
	}

	t.Log(stats.State)
	t.Log(stats.Live.DownloadSpeed.HumanReadable)
	t.Log(stats.Live.UploadSpeed.HumanReadable)
}
