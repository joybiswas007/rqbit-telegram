package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joybiswas007/rqbit-telegram/bot/botutils"
	"github.com/joybiswas007/rqbit-telegram/utils"
)

// Stats returns server-related stats to the user.
func Stats(b *gotgbot.Bot, ctx *ext.Context) error {
	isOwner, errMsg := botutils.SudoChecker(ctx.EffectiveUser.Id, ctx.EffectiveChat.Id)
	if !isOwner {
		botutils.ReplyMessage(b, ctx, fmt.Sprintf("<b>%s</b>", errMsg))
		return nil
	}

	totalDisk, freeDisk, usedDisk, err := utils.DiskUsage()
	if err != nil {
		log.Printf("Failed to get disk usage: %v", err)
		botutils.ReplyMessage(b, ctx, "<b>Failed to get disk usage</b>")
		return nil
	}

	totalMem, freeMem, usedMem := utils.MemoryUsage()
	totalRAM, freeRAM, usedRAM := utils.RAMUsage()

	uptime, err := utils.Uptime()
	if err != nil {
		log.Printf("Failed to get uptime: %v", err)
		botutils.ReplyMessage(b, ctx, "<b>Failed to get uptime</b>")
		return nil
	}

	message := formatStats(totalDisk, freeDisk, usedDisk, totalMem, freeMem, usedMem, totalRAM, freeRAM, usedRAM, uptime)
	botutils.ReplyMessage(b, ctx, message)
	return nil
}

func formatStats(totalDisk, freeDisk, usedDisk, totalMem, freeMem, usedMem, totalRAM, freeRAM, usedRAM uint64, uptime time.Duration) string {
	return fmt.Sprintf(
		"<b>Server Stats:</b>\n"+
			"Disk Usage: %s / %s (used: %s)\n"+
			"Memory Usage: %s / %s (used: %s)\n"+
			"Ram Usage: %s / %s (used: %s)\n"+
			"Uptime: %s",
		utils.HumanReadableBytes(usedDisk), utils.HumanReadableBytes(totalDisk), utils.HumanReadableBytes(freeDisk),
		utils.HumanReadableBytes(usedMem), utils.HumanReadableBytes(totalMem), utils.HumanReadableBytes(freeMem),
		utils.HumanReadableBytes(usedRAM), utils.HumanReadableBytes(totalRAM), utils.HumanReadableBytes(freeRAM),
		uptime.String(),
	)
}
