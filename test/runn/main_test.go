package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/blck-snwmn/example-go/test/runn/api"
	"github.com/go-chi/chi/v5"
	"github.com/k1LoW/runn"
)

var (
	serverURL string
)

func TestMain(m *testing.M) {
	ts := httptest.NewServer(api.HandlerFromMux(NewServer(), chi.NewMux()))
	defer ts.Close()

	serverURL = ts.URL

	m.Run()
}

func TestRunn(t *testing.T) {
	opts := []runn.Option{
		runn.T(t),
		runn.Book("books/example-ownserver.yaml"),
		runn.Runner("req", serverURL),
	}
	o, err := runn.New(opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestRunnN(t *testing.T) {
	// run sequential
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("req", serverURL),
	}
	o, err := runn.Load("./books/example-o*.yaml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestRunnConcurrentN(t *testing.T) {
	// run concurrently
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("req", serverURL),
		runn.RunConcurrent(true, 10), // 10 is magic number
	}
	o, err := runn.Load("./books/example-o*.yaml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(context.Background()); err != nil {
		t.Fatal(err)
	}
}
