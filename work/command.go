package work

import (
	"encoding/json"
	"time"

	"github.com/dmars8047/brolib/chat"
)

const (
	// The name of the queue where work command/tasks can be published so they can be processed by consumers.
	BROCHAT_WORK_QUEUE_NAME = "brochat_work_queue"
)

type CommandType string

const (
	// A command that instructs the recipient to perform post room creation processing.
	COMMAND_TYPE_PROCESS_ROOM_CREATION CommandType = "brochat:command_type:process_room_creation"
	// A command that instructs the recipient to perform post friend request acceptance processing.
	COMMAND_TYPE_PROCESS_FRIEND_REQUEST_ACCEPTANCE CommandType = "brochat:command_type:process_friend_request_acceptance"
	// A command that instructs the recipient to perform post room deletion processing.
	COMMAND_TYPE_PROCESS_ROOM_DELETION CommandType = "brochat:command_type:process_room_deletion"
	// A command that instructs the recipient to save a chat message.
	COMMAND_TYPE_SAVE_CHAT_MESSAGE CommandType = "brochat:command_type:save_chat_message"
	// A command that instructs the recipient to process a users online status change
	COMMAND_TYPE_PROCESS_USER_ONLINE_STATUS_CHANGE CommandType = "brochat:command_type:process_online_status_change"
)

// Acts as an envelope for broadcasted messages
type Command struct {
	// The type of message
	Type CommandType `json:"type"`
	// Content type. Details how the content content should be parsed.
	ContentType string `json:"content_type"`
	// The message data
	Content []byte `json:"content"`
}

// Creates a new FeedMessage. Sets the content as marshaled json bytes and sets the appropriate JSON content type.
func NewCommand(commandType CommandType, command interface{}) (*Command, error) {
	commandBytes, err := json.Marshal(command)

	if err != nil {
		return nil, err
	}

	return &Command{
		ContentType: "application/json",
		Content:     commandBytes,
		Type:        commandType,
	}, nil
}

// Helper function to wrap any struct as a command ready to be published and recieved by consuming workers.
func NewCommandBytes(messageType CommandType, payload interface{}) ([]byte, error) {
	message, err := NewCommand(messageType, payload)

	if err != nil {
		return nil, err
	}

	contentBytes, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	return contentBytes, nil
}

// Contains all necessary data to process a command of type COMMAND_TYPE_SAVE_CHAT_MESSAGE.
type SaveChatMessageCommand struct {
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

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_ROOM_CREATION.
type ProcessRoomCreationCommand struct {
	// The room that was created.
	Room chat.Room `json:"room"`
	// Users to be added to the underlying channel.
	UsersToBeAdded []string `json:"users_to_be_added"`
}

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_FRIEND_REQUEST_ACCEPTANCE.
type ProcessFriendRequestAcceptanceCommand struct {
	// The user that accepted the friend request.
	InitiatingUser chat.UserInfo `json:"initiating_user"`
	// The user that sent the friend request.
	AcceptingUser chat.UserInfo `json:"accepting_user"`
	// The ID of the channel for direct message communication between the users.
	DirectMessageChannel string `json:"direct_message_channel"`
}

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_ROOM_DELETION.
type ProcessRoomDeletionCommand struct {
	// The ID of the room that was deleted.
	RoomId string `json:"room_id"`
}

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_USER_ONLINE_STATUS_CHANGE.
type ProcessUserOnlineStatusChange struct {
	// The ID of the user that changed status.
	UserId string `json:"user_id"`
	// The new status of the user.
	IsOnline bool `json:"is_online"`
	// The time the event occured
	TimeStamp time.Time `json:"timestamp"`
}
