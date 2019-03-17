package main

import (
	"fmt"
	"net/http"

	"github.com/dongri/web-app-on-docker/go/app"
)

const (
	defaultPort = "3001"
)

func main() {
	port := defaultPort

	rooter := app.NewRootHandler()

	// Launch a server instance.
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, rooter)
}
