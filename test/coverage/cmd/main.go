package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blck-snwmn/example-go/test/coverage/server"
)

func main() {
	sv := http.Server{
		Addr:              ":8080",
		Handler:           server.GenHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	fmt.Println("Server is running on", sv.Addr)
	sv.ListenAndServe() //nolint: errcheck
}
