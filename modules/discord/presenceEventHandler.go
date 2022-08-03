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

	member, err := s.GuildMember(m.GuildID, m.User.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if member.User.Bot {
		return
	}

	nickname := member.Nick
	if nickname == "" {
		nickname = member.User.Username
	}

	prevStatus, found := prevStatusMap[m.User.ID]

	if statusIsOnline(m.Status) && (!found || statusIsOffline(prevStatus)) {
		if messageCooldownHandler.CanTriggerPresenceOnline(m.User.ID) {
			s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, fmt.Sprintf("Păzea că vine %s :scream:", nickname))
			messageCooldownHandler.TriggerPresenceOnline(m.User.ID)
		}
	} else if statusIsOffline(m.Status) && (!found || statusIsOnline(prevStatus)) {
		if messageCooldownHandler.CanTriggerPresenceOffline(m.User.ID) {
			s.ChannelMessageSend(spec.Cfg.DiscordCfg.GuildMainChannelId, fmt.Sprintf("Doamne-ajută, a plecat %s :pray:", nickname))
			messageCooldownHandler.TriggerPresenceOffline(m.User.ID)
		}
	}

	prevStatusMap[m.User.ID] = m.Status
}

func statusIsOnline(status discordgo.Status) bool {
	return status == discordgo.StatusOnline || status == discordgo.StatusDoNotDisturb || status == discordgo.StatusIdle
}

func statusIsOffline(status discordgo.Status) bool {
	return status == discordgo.StatusOffline || status == discordgo.StatusInvisible
}
