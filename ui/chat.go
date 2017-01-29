package ui

import "context"

// Talker sends message to target with any options;
// msg as a string should be supported by most backends
// and useful even sending alone, while target and
// options are backend-specific.
type Talker interface {
	Talk(target, msg string, options interface{}) error
}

// ChatMsgHandler handles incoming messages from users;
// each message should invoke ChatMsgHandler
// in a separate goroutine.
type ChatMsgHandler func(ctx context.Context, msg interface{})

// Listener keeps listening to incoming messages,
// do any pre-processing if needed and calls
// ChatMsgHandler with the messages; it will not
// return utils error happens.
type Listener interface {
	Listen(ChatMsgHandler) error
}

// Chatter combines Talker and Listener.
type Chatter interface {
	Talker
	Listener
}
