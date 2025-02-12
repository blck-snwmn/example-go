package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/chi/gen"
	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/chi/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	svi := server.NewServer()

	r := chi.NewRouter()

	gen.HandlerFromMux(svi, r)

	sv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("listening on %s", sv.Addr)
	log.Fatal(sv.ListenAndServe())
}
