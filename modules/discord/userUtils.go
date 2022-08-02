package discord

import (
	"github.com/craciunoiuc/discord-bot/spec"
	"golang.org/x/exp/slices"
)

func userIsCringeMaster(userId string) bool {
	return slices.Contains(spec.Cfg.DiscordCfg.CringeMasterUserIds, userId)
}
