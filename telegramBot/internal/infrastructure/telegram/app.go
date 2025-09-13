package telegram

import (
	"context"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"telegramBot/internal/domain/service"
	"telegramBot/internal/infrastructure/telegram/handlers"
	"telegramBot/pkg/connectors"
	"time"
)

type App struct {
	bot            *telego.Bot
	handlerFactory *handlers.HandlerFactory
}

func NewApp(ctx context.Context, token string, commandService *service.CommandService) *App {
	bot, err := connectors.NewBot(ctx, token)
	if err != nil {
		panic(err)
	}
	options := []telego.LongPollingOption{
		telego.WithLongPollingUpdateInterval(time.Second),
		telego.WithLongPollingRetryTimeout(3 * time.Second),
	}
	updates, err := bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{Timeout: 30}, options...)
	botHandler, err := telegohandler.NewBotHandler(bot, updates)
	if err != nil {
		panic(err)
	}
	handlerFactory := handlers.NewHandlerFactory(botHandler, handlers.NewCommandHandler(ctx, commandService, bot))
	return &App{bot: bot, handlerFactory: handlerFactory}
}

func (a *App) Start() {
	a.handlerFactory.RegisterAllHandlers()
	a.handlerFactory.StartHandler()
}
