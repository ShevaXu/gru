package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

const TokenEnv = "BOT_TOKEN"

func main() {
	// Obtain token from environment variable
	token := os.Getenv(TokenEnv)
	if token == "" {
		panic("No token, please set BOT_TOKEN")
	}
	// When it gets serious
	//os.Unsetenv(TokenEnv)

	api := slack.New(token)
	// used for debug
	//api.SetDebug(true)

	// real-time message
	rtm := api.NewRTM()
	// call once
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch data := msg.Data.(type) {
		// Exist when auth fails
		case *slack.InvalidAuthEvent:
			panic("Invalid credentials!")
		// The client has successfully connected to the server.
		case *slack.HelloEvent:
			fmt.Println("hello")
		// Pings & pongs are already handled by the websocket.
		case *slack.ConnectedEvent:
			fmt.Println("Connected:")
		// A team member's presence changed
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", data)
		// The main payload
		case *slack.MessageEvent:
			fmt.Printf("Message: %s\n", data.Msg.Text)
			// dummy echo-back to the same channel
			rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Echo: %s", data.Msg.Text), data.Channel))
		// TODO: Other useful event types
		//case *slack.LatencyReport, *slack.RTMError:
		default:
			// Ignore other events..
		}
	}
}
