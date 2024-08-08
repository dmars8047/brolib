package chat

import "time"

// A ChatMessage represents a text message sent in to chat channel.
type ChatMessage struct {
	// The Id of the message.
	Id string `json:"id"`
	// The ID of the channel that the message was sent in.
	ChannelId string `json:"channel_id"`
	// The ID of the user that sent the message.
	SenderUserId string `json:"sender_user_id"`
	// The content of the message.
	Content string `json:"content"`
	// The time that the message was sent.
	RecievedAtUtc time.Time `json:"recieved_at_utc"`
}

type UserRelationship struct {
	// The id of the user that the relationship is with.
	UserId string `json:"user_id"`
	// The type of relationship.
	Type RelationshipType `json:"type"`
	// Direct Message Channel Id
	DirectMessageChannelId string `json:"direct_message_channel_id"`
	// Username of the user the relationship is with
	Username string `json:"username"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
	// IsOnline is true if the user is online
	IsOnline bool `json:"is_online"`
}

type ChannelUser struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username.
	Username string `json:"username"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
	// A flag indicating if the user has been removed from the channel for whatever reason (they left, were kicked, banned, deleted, etc) but they remain in the manifest for historical purposes.
	Archived bool `json:"archived"`
}

// A Channel represents a communication channel between two or more users.
type Channel struct {
	// The Id of the channel.
	Id string `json:"id"`
	// The type of the channel.
	Type ChannelType `json:"type"`
	// The users that are members of the channel. This is a list of user info.
	Users []ChannelUser `json:"users"`
}

type InviteUserToRoomRequest struct {
	// The ID of the room
	RoomId string `json:"room_id"`
	// The ID of the user to invite
	UserId string `json:"user_id"`
}

type AcceptRoomInviteRequest struct {
	// The ID of the room
	RoomId string `json:"room_id"`
}

type SendFriendRequestRequest struct {
	// The ID of the user that the friend request is being sent to.
	RequestedUserId string `json:"requested_user_id"`
}

type AcceptFriendRequestRequest struct {
	// The ID of the user that sent the friend request.
	InitiatingUserId string `json:"initiating_user_id"`
}

type RejectFriendRequestRequest struct {
	// The ID of the user that sent the friend request.
	InitiatingUserId string `json:"initiating_user_id"`
}

type UnfriendRequest struct {
	// The ID of the user that is being unfriended.
	UserId string `json:"user_id"`
}

type CancelFriendRequestRequest struct {
	// The ID of the user that the friend request was originally sent to.
	TargetUserId string `json:"target_user_id"`
}
