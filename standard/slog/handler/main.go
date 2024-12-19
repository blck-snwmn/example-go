package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type userIDkey struct{}

func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDkey{}, userID)
}

func getUserID(ctx context.Context) string {
	if v := ctx.Value(userIDkey{}); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

var _ slog.Handler = (*handler)(nil)

type handler struct {
	h slog.Handler
}

// Enabled implements slog.Handler.
func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

// Handle implements slog.Handler.
func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	userID := getUserID(ctx)
	if userID != "" {
		r.AddAttrs(slog.String("user_id", userID))
	}
	return h.h.Handle(ctx, r)
}

// WithAttrs implements slog.Handler.
func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{h: h.h.WithAttrs(attrs)}
}

// WithGroup implements slog.Handler.
func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{h: h.h.WithGroup(name)}
}

func new() *slog.Logger {
	return slog.New(&handler{slog.NewJSONHandler(os.Stdout, nil)})
}

func main() {
	ctx := context.Background()
	log(ctx, slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	log(ctx, new())

	ctx = setUserID(ctx, "123")
	log(ctx, new())
}

func log(ctx context.Context, l *slog.Logger) {
	fmt.Println("========")
	l.InfoContext(ctx, "log 1")

	l = l.With(slog.String("txid", "123"))
	l.InfoContext(ctx, "log 2")

	l = l.WithGroup("group")
	l.InfoContext(ctx, "log 3")

	l = l.With(slog.String("p_id", "456"))
	l.InfoContext(ctx, "log 4")
}
