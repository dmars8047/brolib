package chat

import "time"

type User struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username. This is the same as the username.
	Username string `json:"username"`
	// Profile picture. Its the name of the file as it exists on the media server.
	ProfilePicture string `json:"profile_picture"`
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
	// Profile picture. Its the name of the file as it exists on the media server.
	ProfilePicture string `json:"profile_picture"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
}

type ChannelUser struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username.
	Username string `json:"username"`
	// Profile picture. Its the name of the file as it exists on the media server.
	ProfilePicture string `json:"profile_picture"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
	// A flag indicating if the user has been removed from the channel for whatever reason (they left, were kicked, banned, deleted, etc) but they remain in the manifest for historical purposes.
	Archived bool `json:"archived"`
}
