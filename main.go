package main

import (
	"log"
	"net/http"
	"os"
)

var SNAKE_ID string

func main() {
	http.HandleFunc("/", handleInfo)
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/end", handleEnd)

	SNAKE_ID = os.Getenv("SNAKE_ID")
	if SNAKE_ID == "" {
		log.Fatal("SNAKE_ID environment variable must be set.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
