package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
	"testing/slogtest"
)

func TestSlogHandler(t *testing.T) {
	var buf bytes.Buffer

	slogtest.Run(
		t,
		func(t *testing.T) slog.Handler {
			buf.Reset()
			return &handler{slog.NewJSONHandler(&buf, nil)}
		},
		func(t *testing.T) map[string]any {
			var m map[string]any
			if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
				t.Fatal(err)
			}
			return m
		},
	)
}
