package chat

const (
	GET_USER_URL_SUFFIX              = "/api/brochat/user"
	GET_USERS_URL_SUFFIX             = "/api/brochat/users"
	GET_CHANNEL_URL_SUFFIX           = "/api/brochat/channels/:channelId"
	GET_CHANNEL_MESSAGES_URL_SUFFIX  = "/api/brochat/channels/:channelId/messages"
	SEND_FRIEND_REQUEST_URL_SUFFIX   = "/api/brochat/friends/send-friend-request"
	ACCEPT_FRIEND_REQUEST_URL_SUFFIX = "/api/brochat/friends/accept-friend-request"
	GET_ROOMS_URL_SUFFIX             = "/api/brochat/rooms"
	CREATE_ROOM_URL_SUFFIX           = "/api/brochat/rooms"
	JOIN_ROOM_URL_SUFFIX             = "/api/brochat/rooms/:roomId/join"
)

type RelationshipType uint8

const (
	// This is the default relationship type. It is used when two users are not friends.
	RELATIONSHIP_TYPE_DEFAULT RelationshipType = 1 << iota
	// This relationship type is used when two users are friends.
	RELATIONSHIP_TYPE_FRIEND
	// This relationship type is applied when the user has recieved a friend request from another user.
	RELATIONSHIP_TYPE_FRIEND_REQUEST_RECIEVED
	// This relationship type is applied when the user has sent a friend request to another user.
	RELATIONSHIP_TYPE_FRIENDSHIP_REQUESTED
)

type ChannelType uint8

const (
	// A channel that is used for direct messaging between two users.
	CHANNEL_TYPE_DIRECT_MESSAGE ChannelType = iota
	// A channel that is used for group messages in a room.
	CHANNEL_TYPE_ROOM
)

type RoomMembershipModel string

const (
	// The owner's friends will be allowed to join the room.
	FRIENDS_MEMBERSHIP_MODEL RoomMembershipModel = "friends"
	// The room is public. Anyone can join.
	PUBLIC_MEMBERSHIP_MODEL RoomMembershipModel = "public"
)

type FeedMessageType string

const (
	// Chat message type
	FEED_MESSAGE_TYPE_CHAT_MESSAGE_REQUEST FeedMessageType = "brochat:feed_message_type:chat_message_request"
	// Set active channel message type
	FEED_MESSAGE_TYPE_SET_ACTIVE_CHANNEL_REQUEST FeedMessageType = "brochat:feed_message_type:set_active_channel_request"
	// User online message type
	FEED_MESSAGE_TYPE_USER_ONLINE_EVENT FeedMessageType = "brochat:feed_message_type:user_online_event"
	// User offline message type
	FEED_MESSAGE_TYPE_USER_OFFLINE_EVENT FeedMessageType = "brochat:feed_message_type:user_offline_event"
	// Chat notification message type
	FEED_MESSAGE_TYPE_CHAT_NOTIFICATION FeedMessageType = "brochat:feed_message_type:chat_notification"
	// Chat message message type
	FEED_MESSAGE_TYPE_CHAT_MESSAGE FeedMessageType = "brochat:feed_message_type:chat_message"
	// Friend Request recieved type
	FEED_MESSAGE_TYPE_FRIEND_REQUEST_RECIEVED FeedMessageType = "brochat:feed_message_type:friend_request_recieved"
	// Friend Request accepted type
	FEED_MESSAGE_TYPE_FRIEND_REQUEST_ACCEPTED FeedMessageType = "brochat:feed_message_type:friend_request_accepted"
	// Room created message type
	FEED_MESSAGE_TYPE_ROOM_CREATED FeedMessageType = "brochat:feed_message_type:room_created"
	// User joined a room message type
	FEED_MESSAGE_TYPE_USER_JOINED_ROOM FeedMessageType = "brochat:feed_message_type:user_joined_room"
	// The users profile has been updated. This indicates that the user should refresh their profile in thier local state.
	FEED_MESSAGE_TYPE_USER_PROFILE_UPDATED FeedMessageType = "brochat:feed_message_type:user_profile_updated"
)

type UserProfileUpdateCode uint8

const (
	// The users rooms have been updated.
	USER_PROFILE_UPDATE_CODE_ROOM_UPDATE UserProfileUpdateCode = 0x1
	// The users relationships have been updated.
	USER_PROFILE_UPDATE_REASON_RELATIONSHIP_UPDATE UserProfileUpdateCode = 0x2
)
