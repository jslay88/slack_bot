package main

import (
	"github.com/jslay88/slack_bot/pkg/handlers"
	"github.com/slack-go/slack"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

type Config struct {
	SlackToken string `yaml:"slack_token"`
	HttpPort   string `yaml:"http_port"`
	Debug      bool   `yaml:"debug"`
}

func main() {
	// Build config
	var conf Config
	data, err := os.ReadFile("../../configs/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Create api client
	api := slack.New(
		conf.SlackToken,
		slack.OptionDebug(conf.Debug),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)

	// Fire up HTTP Server
	http.HandleFunc("/healthz", handlers.HealthCheckHandler)
	// Setup Slash Command Handler
	http.HandleFunc("/slack/command", handlers.SlashCommandHandler)
	log.Printf("Starting HTTP Server on port %v...", conf.HttpPort)
	go func() {
		err := http.ListenAndServe(":"+conf.HttpPort, nil)
		if err != nil {
			log.Fatalf("Unable to start HTTP Server. %v", err)
		}
	}()

	// Setup Event Handlers
	handlers.SetupRTMHandlers(api)

	log.Println("Bot started...")
	select {} // Keep the main function alive.
}
