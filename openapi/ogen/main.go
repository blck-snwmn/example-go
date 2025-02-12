package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blck-snwmn/example-go/openapi/ogen/server"
)

func main() {
	sv, err := server.NewServer(server.NewUserRepository())
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr:              ":8080",
		Handler:           sv,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
