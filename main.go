package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to Listen to Server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle Get
		if r.URL.Path == "/foo" {
			// Handle get foo
		}
	}
	if r.Method == http.MethodPost {
		// Handle Post
	}

	w.Write([]byte("Hello, World!"))
}
