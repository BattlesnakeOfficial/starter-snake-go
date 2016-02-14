package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonData map[string]string

func respond(res http.ResponseWriter, obj jsonData) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
}

func handleInfo(res http.ResponseWriter, req *http.Request) {
	respond(res, jsonData{
		"color": "#ff0000",
		"head":  "https://golang.org/doc/gopher/gopherbw.png",
	})
}

func handleStart(res http.ResponseWriter, req *http.Request) {
	respond(res, jsonData{
		"taunt": "battlesnake-go!",
	})
}

func handleMove(res http.ResponseWriter, req *http.Request) {
	data, err := NewMoveRequest(req)
	if err != nil {
		respond(res, jsonData{"move": "north", "taunt": "can't parse this!"})
		return
	}

	snake := data.GetSnake(SNAKE_ID)
	if snake == nil {
		respond(res, jsonData{"move": "north", "taunt": "snake not found!"})
		return
	}
	location := snake.Head()

	log.Printf("Currently at %d,%d", location.X, location.Y)

	board := NewBoard(data)

	chosenDirection := "north"

	directions := []struct {
		name string
		x    int
		y    int
	}{
		{"west", location.X - 1, location.Y},
		{"east", location.X + 1, location.Y},
		{"north", location.X, location.Y - 1},
		{"south", location.X, location.Y + 1},
	}

	for _, direction := range directions {
		tile := board.GetTile(direction.x, direction.y)

		log.Printf("Tile %d,%d is %s", direction.x, direction.y, tile)

		if tile == EMPTY {
			chosenDirection = direction.name
			break
		}
	}

	respond(res, jsonData{
		"move":  chosenDirection,
		"taunt": "go go go!",
	})
}

func handleEnd(res http.ResponseWriter, req *http.Request) {
	respond(res, jsonData{})
}
