package logging

import (
	"log/slog"
	"os"
)

var (
	logger   *slog.Logger
	logLevel = new(slog.LevelVar)
)

func init() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	logger = slog.New(logHandler)
}
