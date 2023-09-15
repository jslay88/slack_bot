package commands

import (
	"github.com/slack-go/slack"
	"log"
)

type GreetCommand struct{}

func (g *GreetCommand) Execute(slashCommand slack.SlashCommand) (string, error) {
	return "Hello", nil
}

func (g *GreetCommand) SlashCommand() string {
	return "/greet"
}

func (g *GreetCommand) RegisterCommand() string {
	cmd := g.SlashCommand()
	log.Printf("Registering Greet Command @ %s...", cmd)
	return cmd
}
