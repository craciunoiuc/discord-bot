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
	"math/rand"
	"strings"

	linq "github.com/ahmetb/go-linq"
	"github.com/bwmarrin/discordgo"
	"github.com/craciunoiuc/discord-bot/modules/markov"
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

func markovGenerate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Trim 'markov' from the start
	msg := m.Content[len("markov "):]

	// Extract the chain name from the message
	words := strings.SplitN(msg, " ", 2)
	chainName := words[0]

	// Check if the chain exists
	if !markov.MarkovChainExists(chainName) {
		s.ChannelMessageSendReply(m.ChannelID, "No markov chain found with name "+chainName+
			". Try these:\n```\n"+markov.MarkovChainList()+"```", m.Reference())
		return
	}

	// Trim the chain name from the message
	msg = msg[len(chainName):]

	// Trim space
	if msg != "" {
		msg = msg[1:]
	}

	// Generate a message
	response := markov.MarkovGenerate(chainName, msg)
	if response == "" {
		s.ChannelMessageSendReply(m.ChannelID, "Could not generate, try something else :(", m.Reference())
	}

	// Send the message
	s.ChannelMessageSendReply(m.ChannelID, response, m.Reference())
}
