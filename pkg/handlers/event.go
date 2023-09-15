package handlers

import (
	"github.com/slack-go/slack"
	"log"
)

func SetupRTMHandlers(api *slack.Client) {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			handleMessageEvent(ev, rtm)
		}
	}
}

func handleMessageEvent(ev *slack.MessageEvent, rtm *slack.RTM) {
	if ev.Channel[0] != 'D' {
		log.Printf("Message received, but was not Direct Message. Ignoring.")
		return
	}
	rtm.SendMessage(rtm.NewOutgoingMessage("I hear you.", ev.Channel))
}
