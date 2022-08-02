package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/craciunoiuc/discord-bot/spec"
)

func PresenceResponse(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	if m.User.ID == s.State.User.ID || !userIsCringeMaster(m.User.ID) {
		return
	}

	if m.Status == discordgo.StatusOnline {
		s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, "Păzea")
	} else if m.Status == discordgo.StatusOffline {
		s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, "Doamne-ajută")
	}
}
