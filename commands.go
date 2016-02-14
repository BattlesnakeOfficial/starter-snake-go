package main

import (
	"encoding/json"
	"net/http"
)

type JSON map[string]string

func respond(res http.ResponseWriter, obj JSON) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
}

func handleInfo(res http.ResponseWriter, req *http.Request) {
	respond(res, JSON{
		"color": "#ff0000",
		"head":  "https://golang.org/doc/gopher/gopherbw.png",
	})
}

func handleStart(res http.ResponseWriter, req *http.Request) {
	respond(res, JSON{
		"taunt": "battlesnake-go!",
	})
}

func handleMove(res http.ResponseWriter, req *http.Request) {
	data, err := NewMoveRequest(req)
	if err != nil {
		respond(res, JSON{"move": "north", "taunt": "can't parse this!"})
		return
	}

	snake := data.GetSnake(SNAKE_ID)
	if snake == nil {
		respond(res, JSON{"move": "north", "taunt": "snake not found!"})
		return
	}

	board := NewBoard(data)
	location := snake.Head()

	var chosenDirection string

	if tile := board.GetTile(location.X, location.Y-1); tile == EMPTY {
		chosenDirection = "north"
	} else if tile := board.GetTile(location.X, location.Y+1); tile == EMPTY {
		chosenDirection = "south"
	} else if tile := board.GetTile(location.X-1, location.Y); tile == EMPTY {
		chosenDirection = "west"
	} else if tile := board.GetTile(location.X+1, location.Y); tile == EMPTY {
		chosenDirection = "east"
	} else {
		chosenDirection = "north"
	}

	respond(res, JSON{
		"move":  chosenDirection,
		"taunt": "go go go!",
	})
}

func handleEnd(res http.ResponseWriter, req *http.Request) {
	respond(res, JSON{})
}
