package main

import (
	"fmt"
	"net/http"

	"github.com/SystemsStuff/ShrinkSync/sampleServer/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", handlers.EchoHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", mux)
}
