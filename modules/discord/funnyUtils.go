package discord

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func handleFunnyMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	normalizationTransform := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	normalizedMessage, _, _ := transform.String(normalizationTransform, m.Content)
	lowerCaseNormalizedMessage := strings.ToLower(normalizedMessage)
	lowerCaseNormalizedMessageWithoutHyphens := strings.ReplaceAll(lowerCaseNormalizedMessage, "-", "")

	if messageIsLoliteas(lowerCaseNormalizedMessageWithoutHyphens) {
		gifUrl := "https://tenor.com/view/frog-frog-dance-rana-ficcata-rana-pazza-gif-23866513"

		message, error := s.ChannelMessageSend(m.ChannelID, gifUrl)
		if message == nil {
			fmt.Println(error.Error())
		}
	}
}

func messageIsLoliteas(message string) bool {
	return strings.Contains(message, "loli") || strings.Contains(message, "lolesc")
}
