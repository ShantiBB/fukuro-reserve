package logger

import (
	"context"
	"encoding"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
)

type Handler struct {
	slog.Handler
	writer io.Writer
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	level := formatLevel(r.Level)
	timeStr := r.Time.Format("2006-01-02 15:04:05")
	msg := r.Message

	var fields []string
	r.Attrs(func(a slog.Attr) bool {
		fields = append(fields, formatAttr(a))
		return true
	})

	output := fmt.Sprintf("%s %s %s %s\n",
		Gray+timeStr+Reset,
		level,
		Cyan+msg+Reset,
		strings.Join(fields, " "),
	)

	_, err := h.writer.Write([]byte(output))
	return err
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithAttrs(attrs),
		writer:  h.writer,
	}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithGroup(name),
		writer:  h.writer,
	}
}

func NewPrettyHandler(w io.Writer, opts *slog.HandlerOptions) *Handler {
	return &Handler{
		Handler: slog.NewTextHandler(w, opts),
		writer:  w,
	}
}

func formatLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return Blue + "DBG" + Reset
	case slog.LevelInfo:
		return Green + "INF" + Reset
	case slog.LevelWarn:
		return Yellow + "WRN" + Reset
	case slog.LevelError:
		return Red + "ERR" + Reset
	default:
		return Cyan + "???" + Reset
	}
}

func formatAttr(attr slog.Attr) string {
	return fmt.Sprintf("%s=%s", Yellow+attr.Key+Reset, formatValue(attr.Value))
}

func formatValue(v slog.Value) string {
	switch v.Kind() {
	case slog.KindString:
		return Green + v.String() + Reset
	case slog.KindInt64:
		return strconv.FormatInt(v.Int64(), 10)
	case slog.KindUint64:
		return strconv.FormatUint(v.Uint64(), 10)
	case slog.KindFloat64:
		return strconv.FormatFloat(v.Float64(), 'f', -1, 64)
	case slog.KindBool:
		return strconv.FormatBool(v.Bool())
	case slog.KindTime:
		return v.Time().Format("2006-01-02 15:04:05.000")
	case slog.KindDuration:
		return v.Duration().String()
	case slog.KindGroup:
		var attrs []string
		for _, a := range v.Group() {
			attrs = append(attrs, formatAttr(a))
		}
		return "{" + strings.Join(attrs, ", ") + "}"
	case slog.KindLogValuer:
		return fmt.Sprintf("%+v", v.Any())
	default:
		if m, ok := v.Any().(encoding.TextMarshaler); ok {
			data, err := m.MarshalText()
			if err != nil {
				return Red + "ERROR: " + err.Error() + Reset
			}
			return Green + string(data) + Reset
		}
		if v.Any() == nil {
			return "null"
		}
		return fmt.Sprintf("%+v", v.Any())
	}
}

func GetTraceID() string {
	return uuid.New().String()
}

func NewTraceID() string {
	return uuid.New().String()
}

func SliceContains(slice []string, str string) bool {
	return slices.Contains(slice, str)
}
