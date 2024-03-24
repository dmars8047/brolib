package chat

import (
	"errors"
	"strings"
)

var (
	ErrMacroTypeUnknown = errors.New("unknown macro type")
)

// Describes a Macros Type.
type MacroType string

const (
	// The Dice Roll Macro.
	MACRO_TYPE_ROLL MacroType = "dice-roll"
	// The Coin Flip Macro.
	MACRO_TYPE_FLIP MacroType = "coin-flip"
	// The Unknown Macro.
	MACRO_TYPE_UNKNOWN MacroType = "unknown"
)

// GetRawMacroType determines if a string represents a Macro request. If it does the type of the macro will be returned.
// If the string does not represent a Macro request, the MACRO_TYPE_UNKNOWN will be returned.
func GetRawMacroType(rawMacro string) MacroType {
	// Get the first word of the message
	val := strings.Split(rawMacro, " ")[0]

	switch val {
	case "/roll":
		return MACRO_TYPE_ROLL
	case "/flip":
		return MACRO_TYPE_FLIP
	default:
		return MACRO_TYPE_UNKNOWN
	}
}

type MacroRequest struct {
	Type MacroType
	Body string
}

var ErrInvalidRollCommandToken = errors.New("invalid roll command token")

type MacroParsingError struct {
	Details string
}

func (parseErr MacroParsingError) Error() string {
	return ErrInvalidRollCommandToken.Error()
}
