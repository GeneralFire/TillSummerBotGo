package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	sq "github.com/Masterminds/squirrel"

	service "github.com/GeneralFire/TillSummerBotGo/internal/botservice"
	"github.com/GeneralFire/TillSummerBotGo/internal/commandhandlers"
	"github.com/GeneralFire/TillSummerBotGo/internal/config"
	"github.com/GeneralFire/TillSummerBotGo/internal/logger"
	"github.com/GeneralFire/TillSummerBotGo/internal/repository"
	"github.com/GeneralFire/TillSummerBotGo/internal/timecalculator"
)

func main() {
	config := GetConfig(".BOT_CONFIG")
	repo, err := repository.New(
		sq.Question,
		config.DB.Driver,
		config.DB.ConnectString,
	)
	if err != nil {
		log.Panic(err)
	}

	loggerInstance := logger.New()
	botService, err := service.Init(
		config.Token,
		loggerInstance,
		repo,
	)
	if err != nil {
		log.Panic(err)
	}

	AddCommandHandlers(botService, repo)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		botService.Stop()
	}()

	botService.ExposeChatButtons()
	err = botService.Serve()
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfig(ymlConfigFile string) *config.BotConfig {
	config, err := config.GetConfigYAML(
		ymlConfigFile,
	)

	if err != nil {
		log.Panic(err)
	}
	return config
}

func AddCommandHandlers(botService *service.BotService, repo service.Repository) {
	timeCalculator := timecalculator.New()
	botService.SetHandler(
		service.CommandDescriptor{
			Prefix: "hello",
			Help:   "Send hello message",
		},
		commandhandlers.GetHelloHandler(&timeCalculator),
	)

	botService.SetHandler(
		service.CommandDescriptor{
			Prefix: "summertime",
			Help:   "Get time till Summer or time passed",
		},
		commandhandlers.GetSummertimeHandler(&timeCalculator),
	)

	botService.SetHandler(
		service.CommandDescriptor{
			Prefix: "subscribe",
			Help:   "Subscibe chat",
		},
		commandhandlers.GetSubscribeHandler(repo),
	)

	botService.SetHandler(
		service.CommandDescriptor{
			Prefix: "unsubscribe",
			Help:   "Unsubscibe chat",
		},
		commandhandlers.GetUnsubscribeHandler(repo),
	)
}
