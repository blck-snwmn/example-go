package main

import (
	"net/http"

	"github.com/blck-snwmn/example-go/openapi/ogen/server"
)

func main() {
	sv, err := server.NewServer(server.NewUserRepository())
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", sv)
}
