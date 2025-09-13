package middleware

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"time"
)

func Logger(next th.Handler) th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		start := time.Now()

		// вызываем основной хендлер
		err := next(ctx, update)

		elapsed := time.Since(start)

		// логируем

		return err
	}
}
