package connectors

import (
	"context"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"log"
)

type BotClient struct {
	Bot     *telego.Bot
	Handler *telegohandler.BotHandler
}

func NewBotClient(token string, ctx context.Context) (*BotClient, error) {
	bot, err := telego.NewBot(token)
	if err != nil {
		return nil, err
	}

	updates, err := bot.UpdatesViaLongPolling(ctx, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	handler, err := telegohandler.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatalln(err)
	}

	return &BotClient{
		Bot:     bot,
		Handler: handler,
	}, nil
}

func (bot *BotClient) Close() error {

}
