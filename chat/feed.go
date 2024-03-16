package chat

import (
	"encoding/json"
)

// Acts as an envelope for broadcasted messages
type FeedMessage struct {
	// The type of message
	Type FeedMessageType `json:"type"`
	// Content type. Details how the content content should be parsed.
	ContentType string `json:"content_type"`
	// The message data
	Content []byte `json:"content"`
}

// Creates a new FeedMessage. Sets the content as marshaled json bytes and sets the appropriate JSON content type.
func NewFeedMessageJSON(messageType FeedMessageType, content interface{}) (*FeedMessage, error) {
	contentBytes, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	return &FeedMessage{
		ContentType: "application/json",
		Content:     contentBytes,
		Type:        messageType,
	}, nil
}

// A notification that a chat message has been recieved.
// Sent to the user when a chat message is recieved but the user is not actively listening to the relvant channel.
type ChatNotification struct {
	// The ID of the channel that the message is being sent in.
	ChannelId string `json:"channel_id"`
}

// Represents an unprocessed chat message.
type ChatMessageRequest struct {
	// The ID of the channel that the message is being sent in.
	ChannelId string `json:"channel_id"`
	// The content of the message.
	Content string `json:"content"`
}

// Describes a Macros Type.
type MacroType string

const (
	// The Dice Roll Macro.
	MACRO_TYPE_ROLL MacroType = "dice-roll"
	// The Coin Flip Macro.
	MACRO_TYPE_FLIP MacroType = "coin-flip"
)

// Represents an unprocessed chat macro.
type ChatMacroRequest struct {
	// The ID of the channel that the message is being sent in.
	ChannelId string `json:"channel_id"`
	// The type of the macro.
	Type MacroType `json:"type"`
	// The macro arguments. The content of the params will be different depending on the type.
	Arguments []string `json:"arguments"`
}

// A request to set the users active channel.
type SetActiveChannelRequest struct {
	// The ID of the channel the user wants to make active.
	ChannelId string `json:"channel_id"`
}

// Represents an event where a user recieves a friend request.
type FriendRequestRecievedEvent struct {
	// The user that sent the friend request.
	InitiatingUser UserInfo `json:"initiating_user"`
	// The user that the friend request was sent to.
	RequestedUser UserInfo `json:"requested_user"`
}

// Represents an event where a user accepts a friend request from another user.
type FriendRequestAcceptedEvent struct {
	// The user that accepted the friend request.
	InitiatingUser UserInfo `json:"initiating_user"`
	// The user that sent the friend request.
	AcceptingUser UserInfo `json:"accepting_user"`
	// The ID of the channel for direct message communication between the users.
	DirectMessageChannel string `json:"direct_message_channel"`
}

// Represents an event when the user's profile has been updated. This indicates that the user should refresh their profile in thier local state.
type UserProfileUpdatedEvent struct {
	// The reason for the update.
	UpdateCode UserProfileUpdateCode `json:"reason"`
}

type ChannelUpdatedEvent struct {
	// The ID of the channel that was updated.
	ChannelId string `json:"channel_id"`
}
