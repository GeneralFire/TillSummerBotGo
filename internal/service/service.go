package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandDescriptor struct {
	Prefix string
	Help   string
}

type HandlerFunc func(tgbotapi.Update) tgbotapi.MessageConfig

type Logger interface {
	Log(string)
}

type BotService struct {
	bot                 *tgbotapi.BotAPI
	handlerMap          map[string]HandlerFunc
	commandsDescriptors []CommandDescriptor
	logger              Logger
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
	descriptor CommandDescriptor,
	handler HandlerFunc,
) {
	d.handlerMap[descriptor.Prefix] = handler
	d.commandsDescriptors = append(d.commandsDescriptors, descriptor)
}

func (d *BotService) Serve() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

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

func (d *BotService) ExposeChatButtons() error {
	commands := make([]tgbotapi.BotCommand, 0, len(d.handlerMap))
	for _, command := range d.commandsDescriptors {
		commands = append(commands,
			tgbotapi.BotCommand{
				Command:     command.Prefix,
				Description: command.Help,
			},
		)
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := d.bot.Request(cfg)
	return err
}

func (d *BotService) Stop() {
	d.bot.StopReceivingUpdates()
}
