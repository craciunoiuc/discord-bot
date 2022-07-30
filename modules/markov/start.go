package markov

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

	"github.com/craciunoiuc/discord-bot/spec"
	"github.com/mb-14/gomarkov"
)

type MarkovData struct {
	Text string `json:"text"`
	Meta string `json:"meta"`
}

type MarkovDataJson struct {
	Data []MarkovData `json:"data"`
}

func MarkovChainExists(name string) bool {
	chain, ok := markovChains[name]
	return ok && chain != nil
}

func MarkovChainList() string {
	result := ""
	for name, _ := range markovChains {
		result += name + "\n"
	}
	return result
}

func MarkovGenerate(name, words string) string {
	order := markovChains[name].Order
	result := ""

	// Split the words into a slice of words
	wordsSlice := strings.FieldsFunc(words, splitSeps)

	// Try to generate the string multiple times
	for retries := 0; result == "" && retries < spec.Cfg.MarkovCfg.MaxRetries; retries++ {
		tokens := make([]string, 0)

		// Generate the first tokens
		for i := 0; i < order; i++ {
			tokens = append(tokens, gomarkov.StartToken)
		}

		// Replace the start tokens with the words provided
		j := len(wordsSlice) - 1
		for i := order - 1; i > 1 && j >= 0; i-- {
			tokens[i] = wordsSlice[j]
			j--
		}

		// Generate one token at a time
		for tokens[len(tokens)-1] != gomarkov.EndToken {
			next, err := markovChains[name].Generate(tokens[(len(tokens) - order):])
			if err != nil {
				fmt.Printf("error generating response for chain %s: %v\n", name, err)
				return ""
			}
			tokens = append(tokens, next)
		}

		// Trim to size the result
		trimSize := order - len(wordsSlice)
		if trimSize < 0 {
			trimSize = 0
		}
		result = strings.Join(tokens[trimSize:len(tokens)-1], " ")

		// Ignore the result if it does not satisfy constraints
		if len(result) < spec.Cfg.MarkovCfg.MinNumberOfChars &&
			len(tokens[order:len(tokens)-1]) < spec.Cfg.MarkovCfg.MinNumberOfWords {
			result = ""
		}
	}

	return result
}

func init() {
	fmt.Printf("Initializing markov chains\n")
	markovChains = make(map[string]*gomarkov.Chain)

	// Create a markov chain for each data provided in the config file
	for _, markovFile := range spec.Cfg.MarkovCfg.Data {
		markovChains[markovFile.Name] = gomarkov.NewChain(spec.Cfg.MarkovCfg.Order)

		markovDataJson := readDataFile(markovFile)

		// Add the data to the chain
		for _, markovData := range markovDataJson.Data {
			if len(markovData.Text) < markovFile.MaxMessageLength &&
				!containsAnyString(markovData.Text, spec.Cfg.MarkovCfg.BlacklistWords) {
				markovChains[markovFile.Name].Add(strings.FieldsFunc(markovData.Text, splitSeps))
			}
		}
	}
}
