package logger

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/lmittmann/tint"
)

type multiHandler struct {
	handlers []slog.Handler
}

func (m multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}

	return false
}

func (m multiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, r)
	}
	return nil
}

func (m multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var newHandler []slog.Handler
	for _, h := range m.handlers {
		newHandler = append(newHandler, h.WithAttrs(attrs))
	}
	return multiHandler{handlers: newHandler}
}

func (m multiHandler) WithGroup(name string) slog.Handler {
	var newHandler []slog.Handler

	for _, h := range m.handlers {
		newHandler = append(newHandler, h.WithGroup(name))
	}

	return multiHandler{handlers: newHandler}
}

func NewLogger(onEnabled bool, fileNameAndPath ...string) {
	logLevel := slog.LevelDebug
	if !onEnabled {
		logLevel = slog.LevelError
	}

	var handlers []slog.Handler

	tintHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      logLevel,
		TimeFormat: "02.01.2006 15:04:05",
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "data" {
				return slog.Attr{}
			}
			return a
		},
	})
	handlers = append(handlers, tintHandler)

	if len(fileNameAndPath) > 0 {
		filePath := fileNameAndPath[0]

		if err := CreateNewLogFile(filePath); err != nil {
			panic(fmt.Sprintf("Failed to create log file: %v", err))
		}

		logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("Failet to ioen log file: %v", err))
		}

		textHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
			Level: logLevel,
		})
		handlers = append(handlers, textHandler)
	}

	slog.SetDefault(slog.New(multiHandler{handlers: handlers}))
}

func Log(v ...interface{}) {
	msg := formatMessage(v...)
	slog.Info(msg)
}

func Warn(v ...interface{}) {
	msg := formatMessage(v...)
	slog.Warn(msg)
}

func Error(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	slog.Error(msg)
	return errors.New(msg)
}

func Debug(v ...interface{}) {
	msg := formatMessage(v...)
	slog.Debug(msg)
}

func formatMessage(v ...interface{}) string {
	var builder strings.Builder
	for _, item := range v {
		builder.WriteString(fmt.Sprintf("%v ", item))
	}

	return strings.TrimSpace(builder.String())
}
