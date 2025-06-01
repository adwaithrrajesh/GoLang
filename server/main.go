package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// hello is a handler for the /api/hello route.
func hello(w http.ResponseWriter, req *http.Request) {
	// Set response header content type
	w.Header().Set("Content-Type", "application/json")

	// Log the request
	log.Printf("Received request: %s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	// Respond with JSON
	response := `{"message": "Hello, this works fine!"}`
	fmt.Fprintln(w, response)
}

func main() {
	// Create a new ServeMux for cleaner route handling
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	// Create a custom server with timeouts and better control
	server := &http.Server{
		Addr:         ":8090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("🚀 Server is starting on http://localhost:8090")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
