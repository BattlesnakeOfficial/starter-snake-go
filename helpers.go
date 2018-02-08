package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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

const LogFormat = `"%s %s %s %d %d" %f`

func LoggingHandler(next http.Handler) http.Handler {
	var logger = log.New(os.Stderr, "", log.LstdFlags)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		loggingWriter := &LoggingResponseWriter{res, http.StatusOK}

		startTime := time.Now()
		next.ServeHTTP(loggingWriter, req)
		elapsedTime := time.Now().Sub(startTime)

		logger.Printf(
			LogFormat,
			req.Method, req.RequestURI, req.Proto,
			loggingWriter.statusCode, 0, elapsedTime.Seconds(),
		)

	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (res *LoggingResponseWriter) WriteHeader(code int) {
	res.statusCode = code
	res.ResponseWriter.WriteHeader(code)
}
