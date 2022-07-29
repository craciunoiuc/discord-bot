package discord

import (
	"fmt"

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

func handleCringeMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if messageIsCringe(m) {
		message, error := s.ChannelMessageSendReply(m.ChannelID, "cringe", m.Reference())
		if message == nil {
			fmt.Println(error.Error())
		}
	}

	if messageIsFromCringeMaster(m) && messageHasStickersFromBlacklistedGuild(s, m) {
		message, error := s.ChannelMessageSendReply(m.ChannelID, "https://tenor.com/view/genshin-impact-zyzz-gif-24919214", m.Reference())
		if message == nil {
			fmt.Println(error.Error())
		}
	}
}
