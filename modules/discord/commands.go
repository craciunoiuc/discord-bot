package discord

import (
	"math/rand"

	linq "github.com/ahmetb/go-linq"
	"github.com/bwmarrin/discordgo"
	spec "github.com/craciunoiuc/discord-bot/spec"
)

func pingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func helpMenu(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "List of commands:\n"
	msg += "```\n"

	for _, key := range commandsCollection.GetSortedKeys() {
		value, found := commandsCollection.Get(key)

		if !found || value.hidden {
			continue
		}

		msg += spec.Cfg.DiscordCfg.Prefix + key + ": " + value.description + "\n"
	}

	msg += "```"

	s.ChannelMessageSend(m.ChannelID, msg)
}

func coinflip(s *discordgo.Session, m *discordgo.MessageCreate) {
	if rand.Intn(2) == 0 {
		s.ChannelMessageSendReply(m.ChannelID, "Head", m.Reference())
	} else {
		s.ChannelMessageSendReply(m.ChannelID, "Tail", m.Reference())
	}
}

func riggedCoinflip(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, "Head", m.Reference())
}

func cringe(s *discordgo.Session, m *discordgo.MessageCreate) {
	var userIds []string

	linq.From(m.Mentions).Where(func(i interface{}) bool {
		return !i.(*discordgo.User).Bot
	}).Select(func(i interface{}) interface{} {
		return i.(*discordgo.User).ID
	}).ToSlice(&userIds)

	cringeObjective = newCringeObjective(userIds)
}

func uncringe(s *discordgo.Session, m *discordgo.MessageCreate) {
	cringeObjective = nil
}
