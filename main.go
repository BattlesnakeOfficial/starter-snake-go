package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Coord `json:"food"`
	Snakes []Snake `json:"snakes"`
}

type Game struct {
	ID string `json:"id"`
}

type StartReqeust struct {
	// TODO
}

type StartResponse struct {
	Color    string `json:"color,omitempty"`
	HeadType string `json:"headType,omitempty"`
	TailType string `json:"tailType,omitempty"`
}

type MoveRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout, omitempty"`
}

type EndRequest struct {
	// TODO
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Your Battlesnake is alive!")
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	response := StartResponse{
		Color:    "#888888",
		HeadType: "regular",
		TailType: "regular",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	response := MoveResponse{
		Move: "right",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	// Nothing to respond with here
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/ping", HandlePing)

	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+8080, nil))
}

// import (
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	http.HandleFunc("/", Index)
// 	http.HandleFunc("/start", Start)
// 	http.HandleFunc("/move", Move)
// 	http.HandleFunc("/end", End)
// 	http.HandleFunc("/ping", Ping)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "9000"
// 	}

// 	// Add filename into logging messages
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)

// 	log.Printf("Running server on port %s...\n", port)
// 	http.ListenAndServe(":"+port, LoggingHandler(http.DefaultServeMux))
// }
