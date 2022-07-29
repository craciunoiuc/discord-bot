package discord

import "github.com/bwmarrin/discordgo"

func getStickerData(s *discordgo.Session, stickerId string) (*discordgo.Sticker, error) {
	endpoint := discordgo.EndpointSticker(stickerId)

	response, error := s.Request("GET", endpoint, nil)
	if error != nil {
		return nil, error
	}

	var sticker discordgo.Sticker

	error = discordgo.Unmarshal(response, &sticker)
	if error != nil {
		return nil, error
	}

	return &sticker, nil
}
