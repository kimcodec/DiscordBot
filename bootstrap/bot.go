package bootstrap

import (
	"github.com/kimcodec/DiscordBot/internal/app"
	"github.com/kimcodec/DiscordBot/internal/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	config, err := viper.InitConfig()
	if err != nil {
		log.Println("failed to init configuration ", err)
	}
	bot := app.NewBot(config)

	err = bot.Run()
	if err != nil {
		log.Println("failed to run bot ", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
