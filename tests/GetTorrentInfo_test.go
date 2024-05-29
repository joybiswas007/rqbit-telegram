package tests

import (
	"testing"

	"github.com/joybiswas007/rqbit-telegram/utils"
)

func TestGetTorrentInfo(t *testing.T) {
	utils.LoadConfig("../config.toml")
	conf := utils.GetConfig()

	info, err := conf.RQBit.GetTorrentInfo(0)
	if err != nil {
		t.Log(err)
	}
	if info == nil {
		t.Log("nothing found")
	}
	t.Log("success")
}
