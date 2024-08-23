package config

import (
	"errors"
)

var (
	errNoWorkingDir = errors.New("no working dir definedd")
	errLoadingConfigFromFile  = errors.New("failed to load config from file")
)
