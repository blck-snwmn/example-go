package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/blck-snwmn/example-go/test/runn/api"
	"github.com/go-chi/chi/v5"
	"github.com/k1LoW/runn"
)

var url string

func TestMain(m *testing.M) {
	srv := NewServer()
	r := chi.NewMux()
	h := api.HandlerFromMux(srv, r)
	ts := httptest.NewServer(h)
	defer ts.Close()

	url = ts.URL

	m.Run()
}

func TestRunn(t *testing.T) {
	opts := []runn.Option{
		runn.T(t),
		runn.Book("books/example-ownserver.yaml"),
		runn.Runner("req", url),
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
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("req", url),
	}
	o, err := runn.Load("./books/example-o*.yaml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(context.Background()); err != nil {
		t.Fatal(err)
	}
}
