package main

import (
    "fmt"
    "net/http"
)


func handleStart(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "start")
}
