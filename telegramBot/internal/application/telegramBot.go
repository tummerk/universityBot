package application

import (
	"context"
	"telegramBot/internal/config"
	"telegramBot/pkg/connectors"
	"telegramBot/pkg/contextx"
)

func Run() {
	ctx, _ := context.WithCancel(context.Background())
	ctx = contextx.WithLogger(ctx, nil)
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	_, err = connectors.NewBotClient(ctx, cfg.TelegramBotToken)
	if err != nil {
		panic(err)
	}
}
