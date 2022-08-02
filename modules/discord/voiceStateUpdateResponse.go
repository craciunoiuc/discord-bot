package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func VoiceStateUpdateResponse(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.UserID == s.State.User.ID || !userIsCringeMaster(m.UserID) {
		return
	}

	if m.SelfStream {
		s.ChannelMessageSend("1002306935653159102", fmt.Sprintf("<@%s> Oprește live-ul!", m.UserID))
	}
}
