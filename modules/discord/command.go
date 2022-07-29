package discord

import "github.com/bwmarrin/discordgo"

type Command struct {
	function    func(s *discordgo.Session, m *discordgo.MessageCreate)
	description string
	hidden      bool
}
