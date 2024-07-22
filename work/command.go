package work

import (
	"time"

	"github.com/dmars8047/brolib/chat"
)

const (
	// The name of the queue where work command/tasks can be published so they can be processed by consumers.
	BROCHAT_WORK_QUEUE_NAME = "brochat_work_queue"
)

type CommandType string

const (
	// A command that instructs the recipient to perform post room deletion processing.
	COMMAND_TYPE_PROCESS_ROOM_DELETION CommandType = "brochat:command_type:process_room_deletion"
	// A command that instructs the recipient to save a chat message.
	COMMAND_TYPE_SAVE_CHAT_MESSAGE CommandType = "brochat:command_type:save_chat_message"
	// A command that instructs the recipient to process a users online status change
	COMMAND_TYPE_PROCESS_USER_ONLINE_STATUS_CHANGE CommandType = "brochat:command_type:process_online_status_change"
	// A command that instructs the recipient to process a change to a channel manifest
	COMMAND_TYPE_PROCESS_CHANNEL_MANIFEST_CHANGE CommandType = "brochat:command_type:process_channel_manifest_change"
)

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

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_ROOM_DELETION.
type ProcessRoomDeletionCommand struct {
	// The room to be deleted.
	Room chat.Room `json:"room"`
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

// Contains all necessary data to process a command of type COMMAND_TYPE_PROCESS_CHANNEL_MANIFEST_CHANGE
type ProcessChannelManifestChange struct {
	// The channel that the change pertains to.
	ChannelId string `json:"channel_id"`
	// The reason/cause of the update.
	Reason string `json:"reason"`
}
