package main

import (
	"net/http"
	"os"
)

const (
	defaultPort = ":8080"
)

func main() {
	port := defaultPort
	if args := os.Args; len(args) > 1 {
		port = args[0]
	}

	http.HandleFunc("/lambda", lambda.LambdaEndpointHandler())
	http.ListenAndServe(port, nil)
}
