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
	"strings"

	"github.com/bwmarrin/discordgo"
	spec "github.com/craciunoiuc/discord-bot/spec"
	"golang.org/x/exp/slices"
)

func messageIsCringe(m *discordgo.MessageCreate) bool {
	return cringeObjective != nil && slices.Contains(cringeObjective.targetUserIds, m.Author.ID)
}

func messageIsFromCringeMaster(m *discordgo.MessageCreate) bool {
	return slices.Contains(spec.Cfg.DiscordCfg.CringeMasterUserIds, m.Author.ID)
}

func guildIsBlacklistedForStickers(guildId string) bool {
	return slices.Contains(spec.Cfg.DiscordCfg.BlacklistStickersGuildIds, guildId)
}

func messageHasStickersFromBlacklistedGuild(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.StickerItems == nil || len(m.StickerItems) == 0 {
		return false
	}

	sticker, error := getStickerData(s, m.StickerItems[0].ID)
	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return guildIsBlacklistedForStickers(sticker.GuildID)
}

func messageHasGenshinRelatedContent(m *discordgo.MessageCreate) bool {
	return strings.Contains(strings.ToLower(m.Content), "genshin")
}

func handleCringeMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if messageIsCringe(m) {
		message, error := s.ChannelMessageSendReply(m.ChannelID, "cringe", m.Reference())
		if message == nil {
			fmt.Println(error.Error())
		}
	}

	if messageIsFromCringeMaster(m) {
		if messageHasStickersFromBlacklistedGuild(s, m) || messageHasGenshinRelatedContent(m) {
			message, error := s.ChannelMessageSendReply(m.ChannelID, "https://tenor.com/view/genshin-impact-zyzz-gif-24919214", m.Reference())
			if message == nil {
				fmt.Println(error.Error())
			}
		} else if strings.Contains(strings.ToLower(m.Content), "lmao") {
			message, error := s.ChannelMessageSendReply(m.ChannelID, "lmao stfu", m.Reference())
			if message == nil {
				fmt.Println(error.Error())
			}
		}
	}
}
