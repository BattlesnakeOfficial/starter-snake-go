package main

import (
    "fmt"
    "net/http"
    "os"
)

func routeRequest(w http.ResponseWriter, r *http.Request) {
    // Ignore favicon for sanity
    if r.URL.Path == "favicon.ico" {
        return
    }

    // Log Request
    fmt.Printf("%s %s\n", r.Method, r.URL)

    // Route
    switch {

    case r.URL.Path == "/" && r.Method == "GET":
        fmt.Fprintf(w, "<a href=\"http://github.com/sendwithus/battlesnake-go\">battlesnake-go</a>")
        return

    case r.URL.Path == "/start" && r.Method == "POST":
        handleStart(w, r)
        return

    case r.URL.Path == "/move" && r.Method == "POST":
        handleMove(w, r)
        return
    case r.URL.Path == "/end" && r.Method == "POST":
        handleEnd(w, r)
        return
    }

    // Method not allowed
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprint(w, "405 Method Not Allowed")
}

func main() {
    // just cause
    http.HandleFunc("/", routeRequest)
    // http.HandleFunc("/favicon.ico", handleNothing)


    // http.HandleFunc("/start", handleStart)
    // http.HandleFunc("/move", handleMove)
    // http.HandleFunc("/end", handleEnd)

    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "9000"
    }

    fmt.Printf("Running server on port %s...\n", port)
    http.ListenAndServe(":" + port, nil)
}
