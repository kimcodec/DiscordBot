package bootstrap

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kimcodec/DiscordBot/handlers"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewBot(t string) {
	dg, err := discordgo.New("Bot " + t)
	if err != nil {
		log.Println("error creating Discord session", err)
		return
	}
	defer dg.Close()

	dg.AddHandler(handlers.MessageCreate)
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if err := dg.Open(); err != nil {
		log.Println("error opening connection", err)
		return
	}
	log.Println("Bot is running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
