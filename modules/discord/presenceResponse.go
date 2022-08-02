package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/craciunoiuc/discord-bot/spec"
)

var prevStatusMap map[string]discordgo.Status

func init() {
	prevStatusMap = make(map[string]discordgo.Status)
}

func PresenceResponse(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	if m.User.ID == s.State.User.ID || !userIsCringeMaster(m.User.ID) {
		return
	}

	prevStatus, found := prevStatusMap[m.User.ID]

	if m.Status == discordgo.StatusOnline {
		if found && prevStatus == discordgo.StatusOnline {
			s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, fmt.Sprintf("<@%s> lasă telefonul!", m.User.ID))
		} else if !found || prevStatus == discordgo.StatusOffline || prevStatus == discordgo.StatusInvisible {
			s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, "Păzea")
		}
	} else if (!found || prevStatus == discordgo.StatusOnline) && m.Status == discordgo.StatusOffline {
		s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, "Doamne-ajută")
	}

	prevStatusMap[m.User.ID] = m.Status
}
