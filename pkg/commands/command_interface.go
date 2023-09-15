package commands

import "github.com/slack-go/slack"

type Command interface {
	Execute(slashCommand slack.SlashCommand) (string, error)
	SlashCommand() string
	RegisterCommand() string
}
