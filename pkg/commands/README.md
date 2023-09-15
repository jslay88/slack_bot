## Slash Commands
To make your own slash command, implement the 
[`Command` interface](command_interface.go) within your own file.

`greet.go`
```go
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
```

You must also register this command on 
[Slack API App Portal](https://api.slack.com/apps/your-app-id/slash-commands)
with the following URL:

    http(s)://your-bot-hostname.com/slack/command

Finally, register the command within 
[`pkg/handlers/command.go`](../handlers/command.go)

```go
func init() {
	// Register commands
	log.Print("Registering Commands...")
	...
	registerCommand(&commands.GreetCommand{})
}
```
