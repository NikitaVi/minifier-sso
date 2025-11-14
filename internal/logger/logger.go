package logger

import (
	"fmt"
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
			NewColorHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		globalLogger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	}
}

func Debug(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	globalLogger.Debug(msg)
}

func Info(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	globalLogger.Info(msg)
}

func Warn(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	globalLogger.Warn(msg)
}

func Error(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	globalLogger.Error(msg)
}

func Fatal(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	globalLogger.Error(msg)
	os.Exit(1)
}
