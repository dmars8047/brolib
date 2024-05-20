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

// Represents an event where a user accepts a friend request from another user.
type FriendRequestAcceptedEvent struct {
	// The user that accepted the friend request.
	InitiatingUser UserInfo `json:"initiating_user"`
	// The user that sent the friend request.
	AcceptingUser UserInfo `json:"accepting_user"`
	// The ID of the channel for direct message communication between the users.
	DirectMessageChannel string `json:"direct_message_channel"`
}

type ChannelUpdatedEvent struct {
	// The ID of the channel that was updated.
	ChannelId string `json:"channel_id"`
}

type UserOnlinStatusChangedEvent struct {
	// The ID of the user that changed status.
	UserId string `json:"user_id"`
	// The new status of the user.
	IsOnline bool `json:"is_online"`
}

// An event that indicates that a room has been deleted.
type RoomDeletedEvent struct {
	// The ID of the room that was deleted.
	RoomId string `json:"room_id"`
}

// ********************
// User Profile Feed Message Events
// ********************

// User profile event for USER_PROFILE_UPDATE_CODE_ROOM_UPDATE
type UserProfileUpdatedEventRoomUpdate struct {
	// The ID of the room that was updated.
	Room Room `json:"room"`
}

// User profile event for USER_PROFILE_UPDATE_CODE_ROOM_ADDED
type UserProfileUpdatedEventRoomAdded struct {
	// The ID of the room that was added.
	Room Room `json:"room"`
}

// User profile event for USER_PROFILE_UPDATE_CODE_ROOM_REMOVED
type UserProfileUpdatedEventRoomRemoved struct {
	// The ID of the room that was removed.
	RoomId string `json:"room_id"`
	// The reason the user was removed from the room.
	Reason string `json:"reason"`
}

// User profile event for USER_PROFILE_UPDATE_REASON_RELATIONSHIP_UPDATE
type UserProfileUpdatedEventRelationshipUpdate struct {
	// The relationship that was updated. In its updated state.
	Relationship UserRelationship `json:"relationship"`
	// The reason the relationship was updated.
	Reason string `json:"reason"`
}

// User profile event for USER_PROFILE_UPDATE_CODE_RELATIONSHIP_ADDED
type UserProfileUpdatedEventRelationshipAdded struct {
	// The new relationship.
	Relationship UserRelationship `json:"relationship"`
	// The reason the relationship was updated.
	Reason string `json:"reason"`
}

// User profile event for USER_PROFILE_UPDATE_CODE_RELATIONSHIP_REMOVED
type UserProfileUpdatedEventRelationshipRemoved struct {
	// The ID of the user who the removed relationship was with.
	UserId string `json:"user_id"`
	// The name of the user who the removed relationship was with.
	Username string `json:"username"`
	// The reason the relationship was removed.
	Reason string `json:"reason"`
}
