package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "DrakeEsdon & HugoKlepsch",
		Color:      "#03fcf4",
		Head:       "pixel",
		Tail:       "pixel",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("START\n")

}

func boarderCheck(you Battlesnake, gameState Board) []string{
    if you.Head.X == 0{//left wall
    	if you.Head.Y == 0{//accounting for the corners (bottom left
			directions := []string{}
			directions = append(directions, "up")
			directions = append(directions, "right")
			return directions
		}
		if you.Head.Y == gameState.Height{//top left corner
			directions := []string{}
			directions = append(directions, "down")
			directions = append(directions, "right")
			return directions
		}
    	directions := []string{}
    	directions = append(directions, "up")
		directions = append(directions, "right")
		directions = append(directions, "down")
    	return directions
	}
	if you.Head.X == gameState.Width{//right wall
		if you.Head.Y == 0{//bottom right corner
			directions := []string{}
			directions = append(directions, "up")
			directions = append(directions, "left")
			return directions
		}
		if you.Head.Y == gameState.Height{//top right corner
			directions := []string{}
			directions = append(directions, "down")
			directions = append(directions, "left")
			return directions
		}
		directions := []string{}
		directions = append(directions, "up")
		directions = append(directions, "right")
		directions = append(directions, "down")
		return directions
	}
	if you.Head.Y == 0{//bottom wall
		directions := []string{}
		directions = append(directions, "up")
		directions = append(directions, "left")
		directions = append(directions, "right")
		return directions
	}
	if you.Head.Y == gameState.Height{//top wall
		directions := []string{}
		directions = append(directions, "down")
		directions = append(directions, "left")
		directions = append(directions, "right")
		return directions
	}
	directions := []string{}
	directions = append(directions, "up")
	directions = append(directions, "right")
	directions = append(directions, "left")
	directions = append(directions, "down")
	return directions
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	var gameState Board = request.Board
	var you Battlesnake = request.You
	availableMoves := boarderCheck(you, gameState)

	move := availableMoves[rand.Intn(len(availableMoves))]

	response := MoveResponse{
		Move: move,
	}

	fmt.Printf("MOVE: %s\n", response.Move)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("END\n")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
