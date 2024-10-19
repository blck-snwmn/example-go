package main

import (
	"log"
	"net/http"

	genoapicodegen "github.com/blck-snwmn/example-go/openapi/mix/gen-oapicodegen"
	"github.com/blck-snwmn/example-go/openapi/mix/server"
)

func main() {
	submux := http.NewServeMux()
	oapiSv := server.NewOapiServer()
	genoapicodegen.HandlerFromMux(oapiSv, submux)

	mainmux := http.NewServeMux()
	ogen, err := server.NewOgenServer()
	if err != nil {
		panic(err)
	}
	mainmux.Handle("/", ogen)
	mainmux.Handle("/", submux)

	log.Printf("listening on %s\n", ":8080")
	http.ListenAndServe(":8080", mainmux)
}
