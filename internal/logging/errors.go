package logging

import (
	"errors"
)

var (
	// ErrInvalidFormat signifies that invalid format text is ised in config.
	// Allowed values are: json, text.
	ErrInvalidFormat = errors.New("provided logging format is invalid")

	// ErrInvalidLevel signifies that invalid log level text is used in the config.
	// Allowed values are: INFO, DEBUG
	ErrInvalidLevel = errors.New("provided logging level is invalid")
)
