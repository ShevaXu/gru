package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

// SlackChat implements Chatter interface. It wraps the slack.Client
// and its RTM websocket for seamless communication.
type SlackChat struct {
	api *slack.Client
	rtm *slack.RTM
}

func (t *SlackChat) Talk(target, msg string, options interface{}) error {
	if options == nil {
		// Sends a simple text-only message to RTM.
		t.rtm.SendMessage(t.rtm.NewOutgoingMessage(msg, target))
		return nil
	}

	switch opt := options.(type) {
	case slack.PostMessageParameters:
		// Post sends a complex message using slack web API
		// (not RTM); its capability is therefor bounded by
		// the API (https://api.slack.com/methods/chat.postMessage).
		_, _, err := t.api.PostMessage(target, msg, opt)
		return err
	case []slack.Attachment:
		params := slack.NewPostMessageParameters()
		params.Attachments = opt
		// the authenticated user will appear as the author of the message,
		// ignoring any values provided for username, icon_url, and icon_emoji
		params.AsUser = true
		_, _, err := t.api.PostMessage(target, msg, params)
		return err
	default:
		fmt.Println("Unhandled options, send message only")
		t.rtm.SendMessage(t.rtm.NewOutgoingMessage(msg, target))
		return nil
	}
}

// slackMsgHandleTimeout is the default timeout
// for handle a slack message
const slackMsgHandleTimeout = 5 * time.Second

// Listen keeps receiving the RTM events and dispatches
// them to the handler function in separate goroutines.
// TODO: logger needed to replace fmt.Println
func (t *SlackChat) Listen(f ChatMsgHandler) error {
	// auth
	auth, err := t.api.AuthTest()
	if err != nil {
		return errors.Wrap(err, "auth test")
	}
	ownId := auth.UserID // bot's ID

	// call once
	go t.rtm.ManageConnection()

	for msg := range t.rtm.IncomingEvents {
		switch data := msg.Data.(type) {
		// Exist when auth fails
		case *slack.InvalidAuthEvent:
			return errors.New("invalid slack credentials")
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
			fmt.Println("Msg:", data.Text)
			// since posting message through web API trigger message-event
			// from Bot's own message, an ID check is necessary to avoid infinite loop
			if data.User != ownId {
				ctx, _ := context.WithTimeout(context.Background(), slackMsgHandleTimeout)
				go f(ctx, data.Msg)
			}
		// TODO: Other useful event types
		//case *slack.LatencyReport, *slack.RTMError:
		default:
			// Ignore other events..
		}
	}

	return nil
}

// NewSlackChat uses the token to set the slack client properly,
// but no connection is made so far.
func NewSlackChat(token string) *SlackChat {
	client := slack.New(token)
	return &SlackChat{
		api: client,
		rtm: client.NewRTM(),
	}
}
