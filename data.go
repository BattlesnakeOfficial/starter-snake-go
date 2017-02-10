package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type GameStartRequest struct {
	GameId string `json:"game_id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type GameStartResponse struct {
	Color   string  `json:"color"`
	HeadUrl *string `json:"head_url,omitempty"`
	Name    string  `json:"name"`
	Taunt   *string `json:"taunt,omitempty"`
}

type MoveRequest struct {
	Food   []Point `json:"food"`
	GameId string  `json:"game_id"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Turn   int     `json:"turn"`
	Snakes []Snake `json:"snakes"`
	You    string  `json:"you"`
}

type MoveResponse struct {
	Move  string  `json:"move"`
	Taunt *string `json:"taunt,omitempty"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	Coords       []Point `json:"coords"`
	HealthPoints int     `json:"health_points"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Taunt        string  `json:"taunt"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

func NewGameStartRequest(req *http.Request) (*GameStartRequest, error) {
	decoded := GameStartRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

func (snake Snake) Head() Point { return snake.Coords[0] }

// Decode [number, number] JSON array into a Point
func (point *Point) UnmarshalJSON(data []byte) error {
	var coords []int
	json.Unmarshal(data, &coords)
	if len(coords) != 2 {
		return errors.New("Bad set of coordinates: " + string(data))
	}
	*point = Point{X: coords[0], Y: coords[1]}
	return nil
}

// Allows decoding a string or number identifier in JSON
// by removing any surrounding quotes and storing in a string
type Identifier string

func (t *Identifier) UnmarshalJSON(data []byte) error {
	*t = Identifier(strings.Trim(string(data), `"`))
	return nil
}
