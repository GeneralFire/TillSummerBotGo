package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GeneralFire/TillSummerBotGo/internal/commandhandlers"
	"github.com/GeneralFire/TillSummerBotGo/internal/config"
	"github.com/GeneralFire/TillSummerBotGo/internal/logger"
	"github.com/GeneralFire/TillSummerBotGo/internal/service"
)

func main() {
	path, _ := os.Getwd()
	log.Println(path)
	botService := GetBotService(".BOT_TOKEN")

	botService.SetHandler(
		service.CommandDescriptor{
			Prefix: "hello",
			Help:   "Send hello message",
		},
		commandhandlers.HelloHandler,
	)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		botService.Stop()
	}()

	botService.ExposeChatButtons()
	err := botService.Serve()
	if err != nil {
		fmt.Println(err)
	}
}

func GetBotService(ymlTokenFile string) service.BotService {
	token, err := config.GetTokenYAML(
		ymlTokenFile,
	)

	if err != nil {
		log.Panic(err)
	}

	loggerInstance := logger.New()

	domainInstance, err := service.Init(
		token,
		loggerInstance,
	)

	if err != nil {
		log.Panic(
			fmt.Errorf("cannot init domain: %w", err),
		)

	}
	return *domainInstance
}