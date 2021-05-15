package commando

import "errors"

var (
	ErrUnknownCommand           = errors.New("unknown command")
	ErrInsufficientArguments    = errors.New("insufficient arguments")
	ErrNameAlreadyUsed          = errors.New("handler is already registered with given name")
	ErrShorthandNameAlreadyUsed = errors.New("handler is already registered with given shorthand name")
)
