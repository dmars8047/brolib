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
	// The No Macro Type. Indicates no macro type.
	MACRO_TYPE_NONE MacroType = "none"
	// The Dice Roll Macro.
	MACRO_TYPE_ROLL MacroType = "dice-roll"
	// The Coin Flip Macro.
	MACRO_TYPE_FLIP MacroType = "coin-flip"
	// The Unknown Macro. Indicates an attempted macro that is not recognized.
	MACRO_TYPE_UNRECOGNIZED MacroType = "unrecognized"
)

// IsMacro determines if a string represents a Macro request. If it does the type of the macro will be returned.
// If the string does not represent a Macro request, the MACRO_TYPE_UNKNOWN will be returned.
func IsMacro(rawMacro string) (bool, MacroType) {
	// Get the first word of the message
	val := strings.Split(rawMacro, " ")[0]

	// Check if the first word is a macro
	if !strings.HasPrefix(val, "/") {
		return false, MACRO_TYPE_NONE
	}

	val = strings.ToLower(val)

	switch val {
	case "/roll":
		return true, MACRO_TYPE_ROLL
	case "/flip":
		return true, MACRO_TYPE_FLIP
	default:
		return true, MACRO_TYPE_UNRECOGNIZED
	}
}

type MacroRequest struct {
	Type MacroType
	Body string
}

type MacroParsingError struct {
	Details string
}
