package connectors

import (
	"context"
	"github.com/mymmrac/telego"
	"telegramBot/pkg/loggerx"
)

func NewBot(ctx context.Context, token string) (*telego.Bot, error) {
	bot, err := telego.NewBot(token, telego.WithDefaultLogger(true, true))
	if err != nil {
		return nil, err
	}
	loggerx.Logger(ctx).Info("telegram created")
	return bot, nil
}
