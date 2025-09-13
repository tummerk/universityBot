package handlers

import (
	"context"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log/slog"
	"telegramBot/internal/domain/service"
	"telegramBot/pkg/loggerx"
)

type CommandHandler struct {
	service *service.CommandService
	bot     *telego.Bot
	ctx     context.Context
}

func NewCommandHandler(ctx context.Context, service *service.CommandService, bot *telego.Bot) *CommandHandler {
	return &CommandHandler{
		service: service,
		bot:     bot,
		ctx:     ctx,
	}
}

func (ch *CommandHandler) HelloWorld(ctx *th.Context, update telego.Update) error {
	text := ch.service.HelloWorld(ctx.Context())
	ch.bot.SendMessage(ctx.Context(), &telego.SendMessageParams{
		ChatID: update.Message.Chat.ChatID(),
		Text:   text,
	})
	loggerx.Logger(ch.ctx).Info("hello world", slog.String("username", update.Message.Chat.Username))
	return nil
}

func (ch *CommandHandler) Start(ctx *th.Context, update telego.Update) error {

	return nil
}
