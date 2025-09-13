package service

import "context"

type CommandService struct {
}

func (bot *CommandService) HelloWorld(ctx context.Context) string {
	return "Hello World"
}
