package tests

import (
	"testing"

	"github.com/joybiswas007/rqbit-telegram/utils"
)

func TestAddTorrent(t *testing.T) {
	utils.LoadConfig("../config.toml")
	conf := utils.GetConfig()

	meta, err := conf.RQBit.AddTorrent("magnet:?xt=urn:btih:")
	if err != nil {
		t.Logf("%v", err)
	}
	t.Log(meta.Details.InfoHash)
	t.Log(meta.Details.Name)
}
