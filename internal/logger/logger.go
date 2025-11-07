package logger

import (
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

const (
	dev  = "local"
	prod = "prod"
)

func Init() {
	cfg := os.Getenv("ENV")

	if cfg == "" {
		panic("env variable ENV is not set")
	}

	switch cfg {
	case dev:
		globalLogger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		globalLogger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	}

}

func Debug(msg string) {
	globalLogger.Debug(msg)
}

func Info(msg string) {
	globalLogger.Info(msg)
}

func Warn(msg string) {
	globalLogger.Warn(msg)
}

func Error(msg string) {
	globalLogger.Error(msg)
}

func Fatal(msg string) {
	globalLogger.Error(msg)
	os.Exit(1)
}
