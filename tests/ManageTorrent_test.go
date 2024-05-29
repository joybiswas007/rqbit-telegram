package tests

import (
	"os"
	"testing"

	"github.com/joybiswas007/rqbit-telegram/utils"
)

func TestManageTorrent(t *testing.T) {
	utils.LoadConfig("../config.toml")
	conf := utils.GetConfig()

	// "start", "pause", "forget", "delete"

	if err := conf.RQBit.ManageTorrent(10, "delete"); err != nil {
		t.Log(err)
		os.Exit(1)
	}

	t.Log("success")
}
