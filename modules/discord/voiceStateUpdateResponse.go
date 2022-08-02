package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/craciunoiuc/discord-bot/spec"
)

func VoiceStateUpdateResponse(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.UserID == s.State.User.ID || !userIsCringeMaster(m.UserID) {
		return
	}

	if m.SelfStream {
		s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, fmt.Sprintf("<@%s> Oprește live-ul!", m.UserID))
	}
}
