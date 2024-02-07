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

type User struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username. This is the same as the username.
	Username string `json:"username"`
	// The users relationships list.
	Relationships []UserRelationship `json:"relationships"`
	// Rooms that the user owns or is a member of
	Rooms []Room `json:"rooms"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
	// CreatedAtUtc is when the user was created
	CreatedAtUtc time.Time `json:"created_at_utc"`
}

type UserInfo struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username.
	Username string `json:"username"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
}

// A Channel represents a communication channel between two or more users.
type Channel struct {
	// The Id of the channel.
	Id string `json:"id"`
	// The type of the channel.
	Type ChannelType `json:"type"`
	// The users that are members of the channel. This is a list of user info.
	Users []UserInfo `json:"users"`
}

type Room struct {
	// The Id of the room
	Id string `json:"id"`
	// The name of the room
	Name string `json:"name"`
	// The rooms channel ID
	ChannelId string `json:"channel_id"`
	// ID of the user who owns the room
	Owner UserInfo `json:"owner"`
	// Membership Model
	MembershipModel RoomMembershipModel `json:"membership_model"`
	// CreatedAtUtc is when the room was created
	CreatedAtUtc time.Time `json:"created_at_utc"`
}

type CreateRoomRequest struct {
	// The name of the room
	Name string `json:"name"`
	// The membership model that the room uses
	MembershipModel string `json:"membership_model"`
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
