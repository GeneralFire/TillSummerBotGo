package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GeneralFire/TillSummerBotGo/internal/commandhandlers"
	"github.com/GeneralFire/TillSummerBotGo/internal/config"
	"github.com/GeneralFire/TillSummerBotGo/internal/logger"
	"github.com/GeneralFire/TillSummerBotGo/internal/service"
)

func main() {
	path, _ := os.Getwd()
	log.Println(path)
	service := GetRawDomain(".BOT_TOKEN")
	service.SetHandler(
		"hello",
		commandhandlers.HelloHandler,
	)

	err := service.Serve()
	fmt.Println(err)
}

func GetRawDomain(ymlTokenFile string) service.BotService {
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
