package viper

import (
	"github.com/kimcodec/DiscordBot/internal/app"
	"github.com/kimcodec/DiscordBot/internal/db"
	"github.com/spf13/viper"
)

func InitConfig() (*app.BotConfig, error) {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	dbURI := viper.GetString("DATABASE_URI")
	botToken := viper.GetString("BOT_TOKEN")
	logLevel := viper.GetString("LOGGER_LEVEL")

	config := app.NewBotConfig(botToken, logLevel, db.NewDBConfig(dbURI))
	return config, nil
}
