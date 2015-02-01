package main

import (
    "fmt"
    "net/http"
)


func handleStart(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,
        `{
            "name": "battlesnake-go",
            "color": "#ff0000",
            "head_url": "http://battlesnake-go.herokuapp.com/",
            "taunt": "battlesnake-go"
        }`)
}
