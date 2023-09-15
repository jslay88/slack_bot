package handlers

import (
	"fmt"
	"github.com/jslay88/slack_bot/pkg/commands"
	"github.com/slack-go/slack"
	"log"
	"net/http"
)

var commandRegistry = make(map[string]commands.Command)

func init() {
	// Register commands
	log.Print("Registering Commands...")
	registerCommand(&commands.GreetCommand{})
}

func registerCommand(cmd commands.Command) {
	commandRegistry[cmd.RegisterCommand()] = cmd
}

func SlashCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received a slash command request: %v", r)
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Printf("Error parsing slash command: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Slash Command parsed. Command: %v", s.Command)

	cmd, ok := commandRegistry[s.Command]
	if !ok {
		log.Printf("Slash command unknown.")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Unknown command.")
		return
	}

	response, cmdErr := cmd.Execute(s)
	if cmdErr != nil {
		log.Printf("Error executing slash command: %v", cmdErr)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, cmdErr)
		return
	}
	fmt.Fprintln(w, response)
}
