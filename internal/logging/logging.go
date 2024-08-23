package logging

import (
	"log/slog"
	"os"
)

const (
	// LogLevelInfo is the default log level.
	LogLevelInfo = "INFO"
	// LogLevelDebug is used to provide detailed info on oprocessing level.
	LogLevelDebug = "DEBUG"
	// TraceLogLevel is used to provide detailed info on records level.

	// LogFormatJSON is a format where all fields are JSON encoded.
	LogFormatJSON = "json"
	// LogFormatText is a human readable format.
	LogFormatText = "text"
)

var Logger *slog.Logger

// Init sets up new logger with a screen output.
func Init(version, level, format string) error {
	var l slog.Level

	// Log level code check and mapping to slog internal values
	switch level {
	case LogLevelDebug:
		l = slog.LevelDebug
	case LogLevelInfo:
		l = slog.LevelInfo
	default:
		return ErrInvalidLevel
	}

	attrs := &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
		Level:       l,
	}

	// Log format code check and creation of specific handler.
	switch format {
	case LogFormatJSON:
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, attrs))
	case LogFormatText:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, attrs))
	default:
		return ErrInvalidFormat
	}

	Logger = Logger.With(
		slog.Group("proc",
			slog.Int("pid", os.Getpid()),
		),
	)

	slog.SetDefault(Logger)

	return nil
}
