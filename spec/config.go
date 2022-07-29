package spec

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
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MarkovConfig struct {
}

type DiscordConfig struct {
	Token               string   `yaml:"token"     env:"DISCORD_TOKEN"     default:""`
	Channel             string   `yaml:"channel"   env:"DISCORD_CHANNEL"   default:""`
	Prefix              string   `yaml:"prefix"    env:"DISCORD_PREFIX"    default:"~"`
	GenshinGuildIds     []string `yaml:"genshinGuildIds"                   default:"[]"`
	CringeMasterUserIds []string `yaml:"cringeMasterUserIds"               default:"[]"`
}

type Config struct {
	DiscordCfg DiscordConfig `yaml:"discord"`
	MarkovCfg  MarkovConfig  `yaml:"markov"`
}

var Cfg Config

func init() {
	fmt.Printf("Loading config file...\n")
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Printf("Read error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}
