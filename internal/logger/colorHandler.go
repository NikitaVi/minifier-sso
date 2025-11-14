package logger

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"log/slog"
	"os"
	"time"
)

type ColorHandler struct {
	opt slog.HandlerOptions
	out *os.File
}

func NewColorHandler(out *os.File, opt *slog.HandlerOptions) *ColorHandler {
	return &ColorHandler{opt: *opt, out: out}
}

func (h *ColorHandler) Handle(_ context.Context, r slog.Record) error {
	timestamp := time.Now().Format("15:04:05")

	var levelColored string

	switch {
	case r.Level == slog.LevelDebug:
		levelColored = color.New(color.FgCyan).Sprint("DEBUG")
	case r.Level == slog.LevelInfo:
		levelColored = color.New(color.FgGreen).Sprint("INFO")
	case r.Level == slog.LevelWarn:
		levelColored = color.New(color.FgYellow).Sprint("WARN")
	case r.Level == slog.LevelError:
		levelColored = color.New(color.FgRed).Sprint("ERROR")
	default:
		levelColored = color.New(color.FgWhite).Sprint(r.Level.String())
	}

	header := color.New(color.FgHiBlack).Sprintf("[%s]", timestamp)

	msg := fmt.Sprintf("%s %s %s\n", header, levelColored, r.Message)

	_, err := h.out.WriteString(msg)
	if err != nil {
		return err
	}

	return nil
}

func (h *ColorHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opt.Level.Level()
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
	return h
}
