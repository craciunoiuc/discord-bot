package main

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
	"os"

	"github.com/spf13/cobra"

	v "github.com/craciunoiuc/discord-bot/internal/version"
)

var (
	version   = "No version provided"
	commit    = "No commit provided"
	buildTime = "No build timestamp provided"

	configFile string
)

// Build the cobra command that handles our command line tool.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "discord-bot -c config.yaml",
		Short: `discord-bot: Multi-purpose Discord bot`,
		Long: `
This is a multi-purpose Discord bot built as a hobby.
It currently supports:
 - markov-chain replies;
 - random image posting;
 - user responses;
`,
		Run:                   doRootCmd,
		DisableFlagsInUseLine: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			showVer, err := cmd.Flags().GetBool("version")
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(0)
			}
			if showVer {
				fmt.Printf(
					"discord-bot %s (%s) built %s\n",
					version,
					commit,
					buildTime,
				)
				os.Exit(0)
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().BoolP(
		"version",
		"V",
		false,
		"Show version and quit",
	)

	rootCmd.PersistentFlags().StringVarP(
		&configFile,
		"config",
		"c",
		"config.yaml",
		"config file",
	)

	// Subcommands
	rootCmd.AddCommand(startCmd)

	return rootCmd
}

// doRootCmd starts the main system
func doRootCmd(cmd *cobra.Command, args []string) {

	fmt.Printf("    ,---,                                                                          ,---,.               ___     \n")
	fmt.Printf("  .'  .' `\\    ,--,                                               ,---,          ,'  .'  \\            ,--.'|_   \n")
	fmt.Printf(",---.'     \\ ,--.'|                           ,---.    __  ,-.  ,---.'|        ,---.' .' |   ,---.    |  | :,'  \n")
	fmt.Printf("|   |  .`\\  ||  |,      .--.--.              '   ,'\\ ,' ,'/ /|  |   | :        |   |  |: |  '   ,'\\   :  : ' :  \n")
	fmt.Printf(":   : |  '  |`--'_     /  /    '     ,---.  /   /   |'  | |' |  |   | |        :   :  :  / /   /   |.;__,'  /   \n")
	fmt.Printf("|   ' '  ;  :,' ,'|   |  :  /`./    /     \\.   ; ,. :|  |   ,',--.__| |        :   |    ; .   ; ,. :|  |   |    \n")
	fmt.Printf("'   | ;  .  |'  | |   |  :  ;_     /    / ''   | |: :'  :  / /   ,'   |        |   :     \\'   | |: ::__,'| :    \n")
	fmt.Printf("|   | :  |  '|  | :    \\  \\    `. .    ' / '   | .; :|  | ' .   '  /  |        |   |   . |'   | .; :  '  : |__  \n")
	fmt.Printf("'   : | /  ; '  : |__   `----.   '   ; :__|   :    |;  : | '   ; |:  |        '   :  '; ||   :    |  |  | '.'| \n")
	fmt.Printf("|   | '` ,/  |  | '.'| /  /`--'  /'   | '.'|\\   \\  / |  , ; |   | '/  '        |   |  | ;  \\   \\  /   ;  :    ; \n")
	fmt.Printf(";   :  .'    ;  :    ;'--'.     / |   :    : `----'   ---'  |   :    :|        |   :   /    `----'    |  ,   /  \n")
	fmt.Printf("|   ,.'      |  ,   /   `--'---'   \\   \\  /                  \\   \\  /          |   | ,'                ---`-'   \n")
	fmt.Printf("'---'         ---`-'                `----'                    `----'           `----'                           \n")
	fmt.Printf("                                                                                                                \n")
	fmt.Printf(" %s\n", v.String())
}

func main() {
	v.SetVersion(&v.Version{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	})

	cmd := NewRootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
