package chat

type BroChatError struct {
	// The error code
	Code BroChatResponseCode `json:"error_code"`
	// The error message
	ErrorDetails []string `json:"error_details"`
}

// NewErrorResponse creates an ErrorResponse with the given code and message.
// Usage: NewErrorResponse(0, "An error occured during validation")
func NewErrorResponse(code BroChatResponseCode, message string) *BroChatError {
	return &BroChatError{
		Code:         code,
		ErrorDetails: []string{message},
	}
}

// NewErrorResponseWithDetails creates an ErrorResponse with the given code and message.
// Usage: NewErrorResponseWithDetails(0, "An error occured during validation", "The username field is required")
func NewErrorResponseWithDetails(code BroChatResponseCode, details ...string) *BroChatError {
	return &BroChatError{
		Code:         code,
		ErrorDetails: details,
	}
}
