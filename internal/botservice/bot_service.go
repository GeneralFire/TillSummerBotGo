package service

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron"
)

type CommandDescriptor struct {
	Command string
	Help    string
}

type HandlerFunc func(tgbotapi.Update) tgbotapi.MessageConfig

type Logger interface {
	Log(string)
}

type Repository interface {
	GetAllSubscribedChat() ([]int64, error)
	SubscribeChat(id int64) error
	UnsubscribeChat(id int64) error
}

type BotService struct {
	bot                 *tgbotapi.BotAPI
	handlerMap          map[string]HandlerFunc
	commandsDescriptors []CommandDescriptor
	logger              Logger
	repository          Repository
	cronInstance        *cron.Cron
}

func Init(
	apiToken string,
	logger Logger,
	repo Repository,
) (*BotService, error) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, err
	}

	cronInstance := cron.New()
	cronInstance.Start()

	return &BotService{
		bot:          bot,
		logger:       logger,
		handlerMap:   make(map[string]HandlerFunc),
		repository:   repo,
		cronInstance: cronInstance,
	}, nil
}

func (d *BotService) SetHandler(
	descriptor CommandDescriptor,
	handler HandlerFunc,
) {
	d.handlerMap[descriptor.Command] = handler
	d.commandsDescriptors = append(d.commandsDescriptors, descriptor)
}

func (d *BotService) CronCallHandlerForAllChat(cronTime, handler string) {
	if _, ok := d.handlerMap[handler]; !ok {
		// TODO: remove panic
		log.Panicf(fmt.Sprintf("cannot find handler for %s", handler))
	}

	updateStub := tgbotapi.Update{Message: &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}}

	if err := d.cronInstance.AddFunc(
		cronTime,
		func() {
			msgText := d.handlerMap[handler](updateStub).Text

			allChats, err := d.repository.GetAllSubscribedChat()
			if err != nil {
				d.logger.Log(
					fmt.Sprintf("get all chats err: %s", err.Error()),
				)
			}
			if len(allChats) == 0 {
				d.logger.Log("no chats")
			}

			for _, chatId := range allChats {
				msg := tgbotapi.NewMessage(chatId, msgText)
				_, err := d.bot.Send(msg)
				if err != nil {
					d.logger.Log(
						fmt.Sprintf(
							"cannot send msg %v to %d", msg, chatId,
						),
					)
				}
			}
		},
	); err != nil {
		log.Panicf(err.Error())
	}
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
				Command:     command.Command,
				Description: command.Help,
			},
		)
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := d.bot.Request(cfg)
	return err
}

func (d *BotService) Stop() {
	d.cronInstance.Stop()
	d.bot.StopReceivingUpdates()
}
