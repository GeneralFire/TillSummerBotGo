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
			Command: "hello",
			Help:    "Send hello message",
		},
		commandhandlers.GetHelloHandler(&timeCalculator),
	)

	summertimeCommandDesc := service.CommandDescriptor{
		Command: "summertime",
		Help:    "Get time till Summer or time passed",
	}

	botService.SetHandler(
		summertimeCommandDesc,
		commandhandlers.GetSummertimeHandler(&timeCalculator),
	)

	botService.SetHandler(
		service.CommandDescriptor{
			Command: "subscribe",
			Help:    "Subscibe chat",
		},
		commandhandlers.GetSubscribeHandler(repo),
	)

	botService.SetHandler(
		service.CommandDescriptor{
			Command: "unsubscribe",
			Help:    "Unsubscibe chat",
		},
		commandhandlers.GetUnsubscribeHandler(repo),
	)

	if err := botService.CronCallHandlerForAllChat(
		"0 0 9 * * *",
		summertimeCommandDesc.Command,
	); err != nil {
		log.Fatal(err)
	}
}
