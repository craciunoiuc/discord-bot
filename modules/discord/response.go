package discord

// SPDX-License-Identifier: BSD-3-Clause
//
// Authors: Cezar Craciunoiu <cezar.craciunoiu@gmail.com>
//			Denis Ehorovici <dehorovici@gmail.com>
//
// Copyright (c) 2022, Universitatea POLITEHNICA Bucharest.  All rights reserved.
// Copyright (c) 2022, Denis Ehorovici. All rights reserved.
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

	"github.com/bwmarrin/discordgo"

	types "github.com/craciunoiuc/discord-bot/internal/types"
	spec "github.com/craciunoiuc/discord-bot/spec"
)

// Collection of all commands
var commandsCollection *types.SortedMap[string, Command]

// Cringe objective used for the cringe command
var cringeObjective *CringeObjective

// Parses all messages sent to the bot and calls the appropriate command
func MessageResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	handleCringeMessages(s, m)

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
		helpMenu(s, m)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())

	commandsCollection = types.NewSortedMap[string, Command]()
	commandsCollection.Add("help", Command{helpMenu, "Displays help menu", false})
	commandsCollection.Add("ping", Command{pingPong, "Pong!", false})
	commandsCollection.Add("markov", Command{markovGenerate, "database [start words] Generate a markov string. Can take additional arguments", false})
	commandsCollection.Add("coinflip", Command{coinflip, "Coinflip", false})
	commandsCollection.Add("cοinflip", Command{riggedCoinflip, "Rigged coinflip", true})
	commandsCollection.Add("cringe", Command{cringe, "Destroy the cringe (mention users)", true})
	commandsCollection.Add("uncringe", Command{uncringe, "Stop cringing... for now", true})
}
