package main

import (
	"context"
	"log/slog"
	"os"
)

type key struct{}

func set(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}

func get(ctx context.Context) *slog.Logger {
	return ctx.Value(key{}).(*slog.Logger)
}

func with(ctx context.Context, args ...any) context.Context {
	l := get(ctx)
	return set(ctx, l.With(args...))
}

func new() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func main() {
	l := new()
	l.Info("Hello, World!")

	ctx := set(context.Background(), l)
	get(ctx).Info("calling doSomething")
	doSomething(ctx)

	ctx = with(ctx, slog.String("id", "123"))
	get(ctx).Info("calling doSomething with id")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	l := get(ctx)
	l.Info("doing something")
}
