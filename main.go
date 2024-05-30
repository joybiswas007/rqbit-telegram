package main

import (
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/joybiswas007/rqbit-telegram/bot/cmd"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

func main() {
	err := utils.LoadConfig("config.toml")
	if err != nil {
		log.Printf("%v", err)
	}

	conf := utils.GetConfig()

	b, err := gotgbot.NewBot(conf.Bot.Token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", cmd.Start))
	dispatcher.AddHandler(handlers.NewCommand("add", cmd.AddTorrent))
	dispatcher.AddHandler(handlers.NewMessage(message.Document, cmd.AddTorrentFile))
	dispatcher.AddHandler(handlers.NewCommand("ids", cmd.GetTorrentIDs))
	dispatcher.AddHandler(handlers.NewCommand("status", cmd.Status))
	dispatcher.AddHandler(handlers.NewCommand("stats", cmd.Stats))
	dispatcher.AddHandler(handlers.NewCommand("pause", cmd.ManageTorrent))
	dispatcher.AddHandler(handlers.NewCommand("resume", cmd.ManageTorrent))
	dispatcher.AddHandler(handlers.NewCommand("remove", cmd.ManageTorrent))
	dispatcher.AddHandler(handlers.NewCommand("delete", cmd.ManageTorrent))

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()

}
