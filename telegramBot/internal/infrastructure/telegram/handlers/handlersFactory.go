package handlers

import (
	th "github.com/mymmrac/telego/telegohandler"
)

// HandlerFactory фабрика для регистрации обработчиков
type HandlerFactory struct {
	botHandler     *th.BotHandler
	commandHandler *CommandHandler
	// другие обработчики
}

// NewHandlerFactory создает новую фабрику обработчиков
func NewHandlerFactory(
	botHandler *th.BotHandler,
	commandHandler *CommandHandler,
) *HandlerFactory {
	return &HandlerFactory{
		botHandler:     botHandler,
		commandHandler: commandHandler,
	}
}

// RegisterAllHandlers регистрирует все обработчики
func (f *HandlerFactory) RegisterAllHandlers() {
	f.registerCommandHandlers()
	// регистрация других обработчиков
}

// RegisterCommandHandlers регистрирует обработчики команд
func (f *HandlerFactory) registerCommandHandlers() {
	// Обработка команды /start
	f.botHandler.Handle(th.Handler(f.commandHandler.Start), th.CommandEqual("start"))
	f.botHandler.Handle(th.Handler(f.commandHandler.HelloWorld), th.CommandEqual("helloWorld"))
}

func (f *HandlerFactory) StartHandler() error {
	return f.botHandler.Start()
}
