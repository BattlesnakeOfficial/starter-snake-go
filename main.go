package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/start", Start)
	http.HandleFunc("/move", Move)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, LoggingHandler(http.DefaultServeMux))
}

const LogFormat = `"%s %d %d" %f`

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		startTime := time.Now()
		next.ServeHTTP(res, req)
		elapsedTime := time.Now().Sub(startTime)

		requestLine := fmt.Sprintf("%s %s %s", req.Method, req.RequestURI, req.Proto)

		log.Printf(
			LogFormat,
			requestLine, http.StatusOK, 0, elapsedTime.Seconds(),
		)

	})
}
