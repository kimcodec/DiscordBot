package main

import (
	"flag"
	"github.com/kimcodec/DiscordBot/bootstrap"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bootstrap.NewBot(Token)
}
