package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kimcodec/DiscordBot/handlers"
	"github.com/kimcodec/DiscordBot/internal/db"
	"github.com/sirupsen/logrus"
	"log"
)

type Bot struct {
	dg     *discordgo.Session
	logger *logrus.Logger
	db     *db.DataBase
	config *BotConfig
}

func NewBot(c *BotConfig) *Bot {
	return &Bot{
		config: c,
		db:     db.NewDB(c.db),
	}
}

func (b *Bot) Run() error {
	dg, err := discordgo.New("Bot " + b.config.token)
	if err != nil {
		log.Println("error opening connection")
		return err
	}

	dg.AddHandler(handlers.MessageCreate)
	dg.AddHandler(handlers.MessageUpdate)
	dg.AddHandler(handlers.MessageDelete)
	dg.AddHandler(handlers.MemberAdd)
	dg.AddHandler(handlers.MemberRemove)
	dg.AddHandler(handlers.MemberUpdate)
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	dg.State.MaxMessageCount = 100

	if err := dg.Open(); err != nil {
		log.Println("error opening connection")
		return err
	}
	b.dg = dg
	log.Println("Bot is running")

	if err := b.db.Open(); err != nil {
		log.Println("error opening database ")
		return err
	}
	log.Println("Database connected successfully")
	logLevel, err := logrus.ParseLevel(b.config.logLevel)
	if err != nil {
		log.Println("failed to parse logger level")
		return err
	}
	b.logger = logrus.New()
	b.logger.SetLevel(logLevel)

	return nil
}
