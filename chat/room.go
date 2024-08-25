package chat

import "time"

type Room struct {
	// The Id of the room
	Id string `json:"id"`
	// The name of the room
	Name string `json:"name"`
	// Image. Its the name of the file as it exists on the media server.
	Image string `json:"image"`
	// The description of the room
	Description string `json:"description"`
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
	// Image. Its the name of the file as it exists on the media server.
	Image string `json:"image"`
	// The description of the room
	Description string `json:"description"`
	// The membership model that the room uses
	MembershipModel string `json:"membership_model"`
}

type UpdateRoomRequest struct {
	// The name of the room
	Name string `json:"name"`
	// The description of the room
	Description string `json:"description"`
	// Image. Its the name of the file as it exists on the media server.
	Image string `json:"image"`
}
