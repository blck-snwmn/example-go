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
	serverURL    string
	badServerURL string
)

func TestMain(m *testing.M) {
	ts := httptest.NewServer(api.HandlerFromMux(NewServer(), chi.NewMux()))
	defer ts.Close()

	serverURL = ts.URL

	badts := httptest.NewServer(api.HandlerFromMux(NewBadServer(), chi.NewMux()))
	defer badts.Close()
	badServerURL = badts.URL

	m.Run()
}

func TestRunn(t *testing.T) {
	t.Parallel()
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

func TestRunnCheckOAPI(t *testing.T) {
	t.Skip("faild test for invalid oapi response")
	opts := []runn.Option{
		runn.T(t),
		runn.Book("books/example-ownserver.yaml"),
		runn.Runner(
			"req",
			badServerURL,
			runn.OpenAPI3("./api/openapi.yaml"),
			runn.SkipValidateRequest(false),
			runn.SkipValidateResponse(false),
		),
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
	t.Parallel()
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
	t.Parallel()
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
