package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	slog.Logger
}

func Init(name string, filePath string) Logger {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		os.Exit(1)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return Logger{*slog.New(handler).With("name", name)}
}
