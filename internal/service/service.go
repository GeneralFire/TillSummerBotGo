package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HandlerFunc func(tgbotapi.Update) tgbotapi.MessageConfig

type Logger interface {
	Log(string)
}

type BotService struct {
	bot        *tgbotapi.BotAPI
	handlerMap map[string]HandlerFunc
	logger     Logger
}

func Init(
	apiToken string,
	logger Logger,
) (*BotService, error) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, err
	}

	return &BotService{
		bot:        bot,
		logger:     logger,
		handlerMap: make(map[string]HandlerFunc),
	}, nil
}

func (d *BotService) SetHandler(
	prefix string,
	handler HandlerFunc,
) {
	d.handlerMap[prefix] = handler
}

func (d *BotService) Serve() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := d.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		command := update.Message.Command()
		handler, ok := d.handlerMap[command]
		if !ok {
			d.logger.Log(
				fmt.Sprintf(
					"Handler for command %s not found!", command,
				),
			)
			continue
		}

		if _, err := d.bot.Send(handler(update)); err != nil {
			return err
		}
	}

	return nil
}
