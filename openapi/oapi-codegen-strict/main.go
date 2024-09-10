package main

import (
	"log"
	"net/http"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen-strict/gen"
	"github.com/blck-snwmn/example-go/openapi/oapi-codegen-strict/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	svi := server.NewServer()

	handler := gen.NewStrictHandler(svi, nil)

	r := chi.NewRouter()

	gen.HandlerFromMux(handler, r)

	sv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Printf("listening on %s", sv.Addr)
	log.Fatal(sv.ListenAndServe())
}
