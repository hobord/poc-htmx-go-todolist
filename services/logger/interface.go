package logger

import "context"

type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

type Attr struct {
	Key   string
	Value any
}

//go:generate mockery --name Logger --structname MockLogger --inpackage --case underscore --disable-version-string
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)

	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)

	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)

	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
}
