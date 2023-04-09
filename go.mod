module github.com/craciunoiuc/discord-bot

go 1.18

require (
	github.com/bwmarrin/discordgo v0.27.0
	github.com/mb-14/gomarkov v0.0.0-20210216094942-a5b484cc0243
	github.com/muesli/termenv v0.14.0
	github.com/spf13/cobra v1.7.0
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	gopkg.in/yaml.v2 v2.4.0
)

require golang.org/x/text v0.7.0

require github.com/aymanbagabas/go-osc52 v1.2.1 // indirect

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
