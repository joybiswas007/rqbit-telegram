package botutils

import "github.com/joybiswas007/rqbit-telegram/utils"

// SudoChecker checks if the command is issued in a private chat by the sudo user.
// It returns true if the conditions are met, otherwise false and an error message if applicable.
func SudoChecker(userId, chatId int64) (bool, string) {
	conf := utils.GetConfig()

	if conf.Bot.SudoID == 0 {
		return false, "sudoId must not be empty."
	}

	if userId != conf.Bot.SudoID {
		return false, "You aren't authorized to use this bot."
	}

	// if chatId != conf.Bot.SudoID {
	// 	return false, "Run this command in private"
	// }
	return true, "You're authorized to run this."
}
