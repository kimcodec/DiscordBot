package app

import "github.com/kimcodec/DiscordBot/internal/db"

type BotConfig struct {
	token    string
	logLevel string
	db       *db.DBConfig
}

func NewBotConfig(t string, ll string, c *db.DBConfig) *BotConfig {
	return &BotConfig{
		token:    t,
		logLevel: ll,
		db:       c,
	}
}
