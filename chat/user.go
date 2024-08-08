package chat

import "time"

type User struct {
	// The user's Id. This is the same as the Id in the idam service.
	Id string `json:"id"`
	// The user's username. This is the same as the username.
	Username string `json:"username"`
	// Profile picture URL
	ProfilePictureUrl string `json:"profile_picture_url"`
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
	// Profile picture URL
	ProfilePictureUrl string `json:"profile_picture_url"`
	// When the user was last online
	LastOnlineUtc time.Time `json:"last_online_utc"`
}
