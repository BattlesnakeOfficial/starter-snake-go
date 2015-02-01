package main

import (
    "fmt"
    "net/http"
)


func handleEnd(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "end")
}
