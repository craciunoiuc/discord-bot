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
	"strings"

	"github.com/bwmarrin/discordgo"
	spec "github.com/craciunoiuc/discord-bot/spec"
)

// Command is a function that handles a command. All command should have the same type.
type Command func(s *discordgo.Session, m *discordgo.MessageCreate)

// The list of commands initialized at startup
var commands map[string]Command

// Parses all messages sent to the bot and calls the appropriate command
func MessageResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't start with the tag
	if !strings.HasPrefix(m.Content, "<@"+s.State.User.ID+"> ") &&
		!strings.HasPrefix(m.Content, spec.Cfg.DiscordCfg.Prefix) {
		return
	}

	idMap := make(map[string]string)
	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	if command, ok := commands[words[0]]; ok {
		command(s, m)
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

	for command := range commands {
		msg += " " + spec.Cfg.DiscordCfg.Prefix + command + "\n"
	}

	msg += "```"

	s.ChannelMessageSend(m.ChannelID, msg)
}

func init() {
	commands = make(map[string]Command)
	commands["ping"] = doCommandPing
	commands["help"] = doCommandHelp
}
