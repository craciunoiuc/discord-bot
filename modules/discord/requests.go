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
