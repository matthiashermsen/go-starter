package logging

import "log/slog"

func SetLevel(level slog.Level) {
	logLevel.Set(level)
}
