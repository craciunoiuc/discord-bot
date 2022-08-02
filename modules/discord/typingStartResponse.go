package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func TypingStartResponse(s *discordgo.Session, m *discordgo.TypingStart) {
	if m.UserID == s.State.User.ID || !userIsCringeMaster(m.UserID) {
		return
	}

	member, err := s.GuildMember(m.GuildID, m.UserID)
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

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ia că scrie %s", nickname))
}
