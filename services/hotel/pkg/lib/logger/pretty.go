package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type SimpleHandler struct {
	out   io.Writer
	level slog.Level
}

func (h *SimpleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *SimpleHandler) Handle(_ context.Context, r slog.Record) error {
	buf := fmt.Sprintf(
		"%s %s msg=%q",
		r.Time.Format("2006-01-02 15:04:05"),
		r.Level.String(),
		r.Message,
	)

	r.Attrs(
		func(a slog.Attr) bool {
			if a.Value.Kind() == slog.KindString {
				buf += fmt.Sprintf(" %s=%q", a.Key, a.Value.String())
			} else {
				buf += fmt.Sprintf(" %s=%v", a.Key, a.Value)
			}
			return true
		},
	)

	buf += "\n"
	_, err := h.out.Write([]byte(buf))
	return err
}

func (h *SimpleHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *SimpleHandler) WithGroup(_ string) slog.Handler {
	return h
}
