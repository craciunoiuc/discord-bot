package discord

// SPDX-License-Identifier: BSD-3-Clause
//
// Authors: Cezar Craciunoiu <cezar.craciunoiu@gmail.com>
//
// Copyright (c) 2022, Universitatea POLITEHNICA Bucharest.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	linq "github.com/ahmetb/go-linq"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slices"

	types "github.com/craciunoiuc/discord-bot/internal/types"
	spec "github.com/craciunoiuc/discord-bot/spec"
)

var commandsCollection *types.SortedMap[string, Command]

var cringeObjective *CringeObjective

func messageIsCringe(m *discordgo.MessageCreate) bool {
	return cringeObjective != nil && slices.Contains(cringeObjective.targetUserIds, m.Author.ID)
}

func messageIsFromCringeMaster(m *discordgo.MessageCreate) bool {
	return slices.Contains(spec.Cfg.DiscordCfg.CringeMasterUserIds, m.Author.ID)
}

func guildIsBlacklistedForStickers(guildId string) bool {
	return slices.Contains(spec.Cfg.DiscordCfg.BlacklistStickersGuildIds, guildId)
}

func messageIsGenshin(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.StickerItems == nil {
		return false
	}

	sticker, error := getStickerData(s, m.StickerItems[0].ID)
	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return guildIsBlacklistedForStickers(sticker.GuildID)
}

// Parses all messages sent to the bot and calls the appropriate command
func MessageResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if messageIsCringe(m) {
		message, error := s.ChannelMessageSendReply(m.ChannelID, "cringe", m.Reference())
		if message == nil {
			fmt.Println(error.Error())
		}
	}

	if messageIsFromCringeMaster(m) && messageIsGenshin(s, m) {
		message, error := s.ChannelMessageSendReply(m.ChannelID, "https://tenor.com/view/genshin-impact-zyzz-gif-24919214", m.Reference())
		if message == nil {
			fmt.Println(error.Error())
		}
	}

	// Ignore all messages that don't start with the tag
	if !strings.HasPrefix(m.Content, "<@"+s.State.User.ID+"> ") &&
		!strings.HasPrefix(m.Content, spec.Cfg.DiscordCfg.Prefix) {
		return
	}

	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}

	idMap := make(map[string]string)
	for _, channel := range channels {
		idMap[channel.Name] = channel.ID
	}

	// Ignore all messages outside of the bot's channel
	if m.ChannelID != idMap[spec.Cfg.DiscordCfg.Channel] {
		return
	}

	// Remove prefix
	if strings.HasPrefix(m.Content, "<@"+s.State.User.ID+"> ") {
		m.Content = strings.TrimPrefix(m.Content, "<@"+s.State.User.ID+"> ")
	} else {
		m.Content = strings.TrimPrefix(m.Content, spec.Cfg.DiscordCfg.Prefix)
	}

	// Split the message into an array of words
	words := strings.Fields(m.Content)

	// Ignore empty messages
	if len(words) == 0 {
		return
	}

	// Get the command and call it
	if command, ok := commandsCollection.Get(words[0]); ok {
		command.function(s, m)
	} else {
		doCommandHelp(s, m)
	}
}

func doCommandPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func doCommandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func doCommandCoinflip(s *discordgo.Session, m *discordgo.MessageCreate) {
	if rand.Intn(2) == 0 {
		s.ChannelMessageSendReply(m.ChannelID, "Head", m.Reference())
	} else {
		s.ChannelMessageSendReply(m.ChannelID, "Tail", m.Reference())
	}
}

func doCommandRiggedCoinflip(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, "Head", m.Reference())
}

func doCommandCringe(s *discordgo.Session, m *discordgo.MessageCreate) {
	var userIds []string

	linq.From(m.Mentions).Where(func(i interface{}) bool {
		return !i.(*discordgo.User).Bot
	}).Select(func(i interface{}) interface{} {
		return i.(*discordgo.User).ID
	}).ToSlice(&userIds)

	cringeObjective = newCringeObjective(userIds)
}

func doCommandUncringe(s *discordgo.Session, m *discordgo.MessageCreate) {
	cringeObjective = nil
}

func init() {
	rand.Seed(time.Now().UnixNano())

	commandsCollection = types.NewSortedMap[string, Command]()
	commandsCollection.Add("help", Command{doCommandHelp, "Displays help menu", false})
	commandsCollection.Add("ping", Command{doCommandPing, "Pong!", false})
	commandsCollection.Add("coinflip", Command{doCommandCoinflip, "Coinflip", false})
	commandsCollection.Add("cοinflip", Command{doCommandRiggedCoinflip, "Rigged coinflip", true})
	commandsCollection.Add("cringe", Command{doCommandCringe, "Destroy the cringe (mention users)", true})
	commandsCollection.Add("uncringe", Command{doCommandUncringe, "Stop cringing... for now", true})
}
