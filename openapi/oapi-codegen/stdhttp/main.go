package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/stdhttp/gen"
	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/stdhttp/server"
)

func main() {
	repository := server.NewUserRepository()
	srv := server.NewServer(repository)

	r := http.NewServeMux()

	gen.HandlerFromMux(srv, r)

	s := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
