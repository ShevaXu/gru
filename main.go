package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"

	"github.com/ShevaXu/gru/nlp/textrazor"
	"github.com/ShevaXu/gru/ui"
	"github.com/ShevaXu/gru/utils"
)

const (
	TokenEnv          = "BOT_TOKEN"
	TextRazorTokenEnv = "TEXTRAZOR_TOKEN"
)

func Echo(chat ui.Chatter) ui.ChatMsgHandler {
	return func(ctx context.Context, msg interface{}) {
		if m, ok := msg.(slack.Msg); ok {
			//params := slack.NewPostMessageParameters()
			//params.AsUser = true
			//err := chat.Talk(m.Channel, fmt.Sprintf("Echo: %s", m.Text), params)
			err := chat.Talk(m.Channel, fmt.Sprintf("Echo: %s", m.Text), nil)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func AnalyzeText(chat ui.Chatter, cl *textrazor.Client) ui.ChatMsgHandler {
	return func(ctx context.Context, msg interface{}) {
		if m, ok := msg.(slack.Msg); ok {
			text := m.Text
			res, err := cl.Query(textrazor.EntitiesWordsRelations, text)
			if err != nil {
				log.Println(err)
				return
			}
			output := res.Resp.Sentences[0].OneLine()
			if len(res.Resp.Entities) > 0 {
				output += (" - " + res.Resp.Entities[0].WikiLink)
			}
			err = chat.Talk(m.Channel, output, nil)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func main() {
	// Obtain token from environment variable
	token := os.Getenv(TokenEnv)
	if token == "" {
		panic("No token, please set BOT_TOKEN")
	}
	// Optional nlp API
	nlpToken := os.Getenv(TextRazorTokenEnv)
	// When it gets serious
	//os.Unsetenv(TokenEnv)
	//os.Unsetenv(TextRazorTokenEnv)

	// init chatter
	lg := log.New(os.Stdout, "[SLACK] ", log.Lshortfile|log.LstdFlags)
	chat := ui.NewSlackChat(token, lg)

	var f ui.ChatMsgHandler
	if nlpToken != "" {
		cl := textrazor.NewClient(nlpToken, utils.DefaultClient)
		f = AnalyzeText(chat, cl)
	} else {
		f = Echo(chat)
	}

	log.Println(chat.Listen(f))
}
