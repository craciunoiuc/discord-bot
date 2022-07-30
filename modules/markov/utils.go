package markov

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/craciunoiuc/discord-bot/spec"
	"github.com/mb-14/gomarkov"
)

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

var markovChains map[string]*gomarkov.Chain

func splitSeps(r rune) bool {
	return strings.ContainsRune(" .?,;:()[]{}\n\t\r", r)
}

func containsAnyString(s string, strs []string) bool {
	for _, str := range strs {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

func readDataFile(markovFile spec.MarkovDataConfig) *MarkovDataJson {
	// Unzip the data path
	archivedData, err := zip.OpenReader(markovFile.Path)
	if err != nil {
		panic(fmt.Errorf("error opening zip file %s: %v", markovFile.Name, err))
	}
	defer archivedData.Close()

	// Archive should have only one file
	if len(archivedData.File) != 1 {
		panic(fmt.Errorf("zip file %s should have only one file", markovFile.Name))
	}

	// Open the file
	file := archivedData.File[0]
	fileReader, err := file.Open()
	if err != nil {
		panic(fmt.Errorf("error opening file %s: %v", markovFile.Name, err))
	}

	// Read the file
	data, err := ioutil.ReadAll(fileReader)
	if err != nil {
		panic(fmt.Errorf("error reading file %s: %v", markovFile.Name, err))
	}

	// Unmarshal the data to json
	markovDataJson := MarkovDataJson{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&markovDataJson)

	return &markovDataJson
}
