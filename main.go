package main

import (
	"log"
	"net/http"

	"github.com/neuronOS/cmd"
	"github.com/neuronOS/controller"
)

func main() {
	// Commander Object
	cmdr := cmd.NewCommander()

	// Mux Router
	mux := http.NewServeMux()

	// Register the "/execute" route with the command handler
	mux.HandleFunc("/execute", controller.HandleCommand(cmdr))

	// Start the HTTP server on port 8080 with the
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
