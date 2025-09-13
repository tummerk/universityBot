package application

import (
	"context"
	"sync"
	"telegramBot/internal/config"
	"telegramBot/internal/domain/service"
	"telegramBot/internal/infrastructure/telegram"
	"telegramBot/pkg/connectors"
	"telegramBot/pkg/contextx"
	"telegramBot/pkg/loggerx"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	slog := &connectors.Slog{
		Debug:    false,
		FileName: "logs.txt",
	}
	ctx = contextx.WithLogger(ctx, slog.Logger())
	cfg, err := config.Load()
	if err != nil {
		loggerx.Logger(ctx).Error("load config fail")
	}
	commandService := service.CommandService{}
	telegramApp := telegram.NewApp(ctx, cfg.TelegramBotToken, &commandService)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		telegramApp.Start()
	}()
	wg.Wait()
}
