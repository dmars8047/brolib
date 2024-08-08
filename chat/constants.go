package chat

const (
	GET_USER_URL_SUFFIX              = "/api/brochat/user"
	GET_USERS_URL_SUFFIX             = "/api/brochat/users"
	GET_CHANNEL_URL_SUFFIX           = "/api/brochat/channels/:channelId"
	GET_CHANNEL_MESSAGES_URL_SUFFIX  = "/api/brochat/channels/:channelId/messages"
	SEND_FRIEND_REQUEST_URL_SUFFIX   = "/api/brochat/friends/send-friend-request"
	ACCEPT_FRIEND_REQUEST_URL_SUFFIX = "/api/brochat/friends/accept-friend-request"
	REJECT_FRIEND_REQUEST_URL_SUFFIX = "/api/brochat/friends/reject-friend-request"
	CANCEL_FRIEND_REQUEST_URL_SUFFIX = "/api/brochat/friends/cancel-friend-request"
	UNFRIEND_USER_URL_SUFFIX         = "/api/brochat/friends/unfriend"
	GET_ROOMS_URL_SUFFIX             = "/api/brochat/rooms"
	CREATE_ROOM_URL_SUFFIX           = "/api/brochat/rooms"
	UPDATE_ROOM_URL_SUFFIX           = "/api/brochat/rooms/:roomId"
	DELETE_ROOM_URL_SUFFIX           = "/api/brochat/rooms/:roomId"
	JOIN_ROOM_URL_SUFFIX             = "/api/brochat/rooms/:roomId/join"
	LEAVE_ROOM_URL_SUFFIX            = "/api/brochat/rooms/:roomId/leave"
	UPLOAD_PROFILE_PICTURE_SUFFIX    = "/api/brochat/media/profile-picture"
	UPLOAD_CHANNEL_FILE_SUFFIX       = "/api/brochat/media/channel/:channelId"
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
	// FRIENDS_MEMBERSHIP_MODEL RoomMembershipModel = "friends_only"
	// The room is public. Anyone can join.
	PUBLIC_MEMBERSHIP_MODEL RoomMembershipModel = "public"
)

type FeedMessageType string

const (
	// Chat message request type. This represents a raw chat message coming from the user.
	FEED_MESSAGE_TYPE_CHAT_MESSAGE_REQUEST FeedMessageType = "brochat:feed_message_type:chat_message_request"
	// Set active channel message type. This represents when a user has navigated to a new channel or away from a channel.
	FEED_MESSAGE_TYPE_SET_ACTIVE_CHANNEL_REQUEST FeedMessageType = "brochat:feed_message_type:set_active_channel_request"
	// Chat notification message type. This is just a notification of the event. It does not contain the content of the chat message.
	FEED_MESSAGE_TYPE_CHAT_NOTIFICATION FeedMessageType = "brochat:feed_message_type:chat_notification"
	// Chat message message type. Represents a processed chat message ready to be shown to users.
	FEED_MESSAGE_TYPE_CHAT_MESSAGE FeedMessageType = "brochat:feed_message_type:chat_message"
	// The feed message indicating that a channel has been updated.
	FEED_MESSAGE_TYPE_CHANNEL_UPDATED FeedMessageType = "brochat:feed_message_type:channel_updated"
	// The feed message that represents a macro request. This is a special kind of chat message that runs some logic to generate/format the chat message content.
	FEED_MESSAGE_TYPE_MACRO_REQUEST FeedMessageType = "brochat:feed_message_type:macro_request"
	// The feed message type that represents an event where a room (or rooms) a user belongs to has been updated.
	FEED_MESSAGE_TYPE_USER_ROOMS_UPDATED FeedMessageType = "brochat:feed_message_type:user_rooms_updated"
	// The feed message type that represents an event where a user has been added to a room (or rooms).
	FEED_MESSAGE_TYPE_USER_ROOMS_ADDED FeedMessageType = "brochat:feed_message_type:user_rooms_added"
	// The feed message type that represents an event where a user has been removed from a room (or rooms).
	FEED_MESSAGE_TYPE_USER_ROOMS_REMOVED FeedMessageType = "brochat:feed_message_type:user_rooms_removed"
	// The feed message type that represents an event where a user relationship (or relationships) has been updated.
	FEED_MESSAGE_TYPE_USER_RELATIONSHIPS_UPDATED FeedMessageType = "brochat:feed_message_type:user_relationships_updated"
	// The feed message type that represents an event where a user relationship (or relationships) has been added.
	FEED_MESSAGE_TYPE_USER_RELATIONSHIPS_ADDED FeedMessageType = "brochat:feed_message_type:user_relationships_added"
	// The feed message type that represents an event where a user relationship (or relationships) has been removed.
	FEED_MESSAGE_TYPE_USER_RELATIONSHIPS_REMOVED FeedMessageType = "brochat:feed_message_type:user_relationships_removed"
	// User online message type. This represents when a user's online status has changed.
	FEED_MESSAGE_TYPE_USER_ONLINE_STATUS_UPDATED_EVENT FeedMessageType = "brochat:feed_message_type:user_online_status_updated_event"
)

// BroChatResponseCode is a numeric representation of the error code returned by the BroChat API.
type BroChatResponseCode uint8

// Server side error codes
const (
	// Indicates an unhandled error.
	BROCHAT_RESPONSE_CODE_UNHANDLED_ERROR BroChatResponseCode = iota
	// Indicates a forbidden operation error. This means the user does not have permission to perform the operation.
	BROCHAT_RESPONSE_CODE_FORBIDDEN_ERROR
	// Indicates a validation error. This means the associated request parameters were invalid.
	BROCHAT_RESPONSE_CODE_VALIDATION_ERROR
	// Indicates a request body parsing error. This means the request body could not be parsed.
	BROCHAT_RESPONSE_CODE_REQUEST_PARSE_ERROR
	// Indicates a not found error. This means the requested resource was not found.
	BROCHAT_RESPONSE_CODE_NOT_FOUND_ERROR
	// Indicates a data conflict error. This means the request could not be completed due to a conflict with the current state of the resource.
	BROCHAT_RESPONSE_CODE_DATA_CONFLICT_ERROR
	// Indicates an invalid operation error. This means the requested operation is invalid. Example: Trying to become friends with yourself.
	BROCHAT_RESPONSE_CODE_INVALID_OPERATION
	// Indicates an unauthorized operation error. This means the user is not authorized to perform the requested operation.
	BROCHAT_RESPONSE_CODE_UNAUTHORIZED_ERROR
)

// Client side error codes
const (
	// Indicates an invalid host address error. This means the address that the client is trying to connect to is invalid.
	BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS = iota + 64
	// Timeout error. This indicates that the BroChat API did not respond in a timely manner.
	BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR
	// Indicates the request content was not formatted properly.
	BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR
	// Indicates that the response from the server was unexpected and could not be parsed.
	BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR
	// Indicates a generic request error.
	BROCHAT_RESPONSE_CODE_GENERIC_REQUEST_ERROR
	// Indicates a generic connection error.
	BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR
)

// Success codes
const (
	// Succese code 128 indicates a successful operation.
	BROCHAT_RESPONSE_CODE_SUCCESS BroChatResponseCode = iota + 128
	// Success code 256 indicates a succesful operation with no content.
	BROCHAT_RESPONSE_CODE_NO_CONTENT
)
