package chat

import (
	"encoding/json"
	"time"
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
func NewFeedMessage(messageType FeedMessageType, content interface{}) (*FeedMessage, error) {
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

// Creates a new FeedMessage. Sets the content as marshaled json bytes and sets the appropriate JSON content type. Returns the bytes of the message ready to be sent over the wire.
// Example Usage: NewFeedMessageBytes(chat.FEED_MESSAGE_TYPE_CHAT_MESSAGE_REQUEST, chat.ChatMessageRequest{ChannelId: "123", Content: "Hello World"})
func NewFeedMessageBytes(messageType FeedMessageType, content interface{}) ([]byte, error) {
	message, err := NewFeedMessage(messageType, content)

	if err != nil {
		return nil, err
	}

	contentBytes, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	return contentBytes, nil
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

type ChannelUpdatedEvent struct {
	// The ID of the channel that was updated.
	ChannelId string `json:"channel_id"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_ROOM_UPDATED
type UserRoomsUpdatedEvent struct {
	// The ID of the room that was updated.
	Rooms []Room `json:"room"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_ROOM_ADDED
type UserRoomsAddedEvent struct {
	// The ID of the room that was added.
	Rooms []Room `json:"room"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_ROOM_REMOVED
type UserRoomsRemovedEvent struct {
	// The ID of the room that was removed.
	RoomIds []string `json:"room_id"`
	// The reason the user was removed from the room.
	Reason string `json:"reason"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_RELATIONSHIP_UPDATED
type UserRelationshipsUpdatedEvent struct {
	// The relationship that was updated. In its updated state.
	Relationships []UserRelationship `json:"relationship"`
	// The reason the relationship was updated.
	Reason string `json:"reason"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_RELATIONSHIP_ADDED
type UserRelationshipsAddedEvent struct {
	// The new relationship.
	Relationships []UserRelationship `json:"relationship"`
	// The reason the relationship was updated.
	Reason string `json:"reason"`
}

// User profile event for FEED_MESSAGE_TYPE_USER_RELATIONSHIP_REMOVED
type UserRelationshipsRemovedEvent struct {
	// The ID of the user who the removed relationship was with.
	UserIds []string `json:"user_id"`
	// The reason the relationship was removed.
	Reason string `json:"reason"`
}

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_USER_ONLINE_STATUS_CHANGE.
type UserOnlineStatusChangeEvent struct {
	// The ID of the user that changed status.
	UserId string `json:"user_id"`
	// The new status of the user.
	IsOnline bool `json:"is_online"`
	// The time the event occured
	TimeStamp time.Time `json:"timestamp"`
}
