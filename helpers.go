package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}

func dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
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
