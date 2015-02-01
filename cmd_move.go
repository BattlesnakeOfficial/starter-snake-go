package main

import (
    "fmt"
    "net/http"
)


func handleMove(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,
        `{
            "move": "up",
            "taunt": "going up!"
        }`)
}
