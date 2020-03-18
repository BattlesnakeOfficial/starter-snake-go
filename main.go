package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
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

type StartRequest struct {
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

// HandleStart is called at the start of each game your Battlesnake is playing.
// The StartRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := StartRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	response := StartResponse{
		Color:    "#888888",
		HeadType: "regular",
		TailType: "regular",
	}

	fmt.Print("START\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the MoveRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := MoveRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]

	response := MoveResponse{
		Move: move,
	}

	fmt.Printf("MOVE: %s\n", response.Move)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := EndRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Nothing to respond with here
	fmt.Print("END\n")
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
