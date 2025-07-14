package bot

import (
	tb "gopkg.in/telebot.v3"
	"mybot/config"
)

func CheckPermission(c tb.Context) bool {
	sender := c.Sender()
	chat := c.Chat()

	userAllowed := len(config.Cfg.AllowedUserIDs) == 0
	for _, id := range config.Cfg.AllowedUserIDs {
		if id == sender.ID {
			userAllowed = true
			break
		}
	}

	groupAllowed := len(config.Cfg.AllowedGroupIDs) == 0
	for _, id := range config.Cfg.AllowedGroupIDs {
		if chat.ID == id {
			groupAllowed = true
			break
		}
	}

	return userAllowed && groupAllowed
}
