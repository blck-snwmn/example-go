package main

import (
	"log"
	"net/http"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/stdhttp/gen"
	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/stdhttp/server"
)

func main() {
	srv := server.NewServer()

	r := http.NewServeMux()

	gen.HandlerFromMux(srv, r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Printf("listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
