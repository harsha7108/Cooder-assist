package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	slog.Logger
}

func Init(name string) Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return Logger{*slog.New(handler).With("name", name)}
}
