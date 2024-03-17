package chat

// BroChatError is the response returned when an error
// is encoutered during the processing of a request to the BroChat API.
type BroChatError struct {
	Code    uint16 `json:"error_code"`
	Message string `json:"error_message"`
}

// Error returns the error message for the BroChatError
func (err BroChatError) Error() string {
	return err.Message
}

// NewErrorResponse creates an ErrorResponse with the given code and message.
// Usage: NewErrorResponse(0, "An error occured during validation")
func NewErrorResponse(code uint16, message string) *BroChatError {
	return &BroChatError{
		Code:    code,
		Message: message,
	}
}

// NewUnhandledErrorResponse creates an ErrorResponse with the code and message for an unhandled error.
func NewUnhandledErrorResponse() *BroChatError {
	return &BroChatError{
		Code:    ERROR_CODE_UNHANDLED,
		Message: "an unhandled/unexpected error occured",
	}
}

const (
	// Error codes
	// Error code 0 indicates an unhandled error. This means there was a server error.
	ERROR_CODE_UNHANDLED = 0x0001
)
