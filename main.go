package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"

	"github.com/ShevaXu/gru/ui"
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

	lg := log.New(os.Stdout, "[SLACK] ", log.Lshortfile|log.LstdFlags)
	chat := ui.NewSlackChat(token, lg)

	fmt.Println(chat.Listen(func(ctx context.Context, msg interface{}) {
		if m, ok := msg.(slack.Msg); ok {
			//params := slack.NewPostMessageParameters()
			//params.AsUser = true
			//err := chat.Talk(m.Channel, fmt.Sprintf("Echo: %s", m.Text), params)
			err := chat.Talk(m.Channel, fmt.Sprintf("Echo: %s", m.Text), nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	}))
}
