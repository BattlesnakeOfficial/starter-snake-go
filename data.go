package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type MoveRequest struct {
	Snakes []Snake    `json:"snakes"`
	Turn   int        `json:"turn"`
	Food   []Point    `json:"food"`
	Height int        `json:"height"`
	Width  int        `json:"width"`
	Game   Identifier `json:"game"`
	Mode   string     `json:"mode"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

func (req *MoveRequest) GetSnake(id string) *Snake {
	for _, snake := range req.Snakes {
		if snake.ID == Identifier(id) {
			return &snake
		}
	}
	return nil
}

type Snake struct {
	ID      Identifier `json:"id"`
	Name    string     `json:"name"`
	Status  string     `json:"status"`
	Coords  []Point    `json:"coords"`
	Kills   int        `json:"kills"`
	Taunt   string     `json:"taunt"`
	Age     int        `json:"age"`
	Health  int        `json:"health"`
	Message string     `json:"message"`
}

func (snake Snake) Head() Point { return snake.Coords[0] }

type Board struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

type Tile int

const (
	EMPTY Tile = iota
	BOUNDARY
	SNAKE_HEAD
	SNAKE_BODY
	FOOD
)

func (t Tile) String() string {
	switch t {
	case EMPTY:
		return "Tile{EMPTY}"
	case BOUNDARY:
		return "Tile{BOUNDARY}"
	case SNAKE_HEAD:
		return "Tile{SNAKE_HEAD}"
	case SNAKE_BODY:
		return "Tile{SNAKE_BODY}"
	case FOOD:
		return "Tile{FOOD}"
	}
	return "Tile{}"
}

func NewBoard(req *MoveRequest) *Board {
	board := Board{
		Width:  req.Width,
		Height: req.Height,
		Tiles:  make([][]Tile, req.Width),
	}

	// Tiles[x][y] = EMPTY
	for x := 0; x < board.Width; x++ {
		row := make([]Tile, board.Height)
		for y := 0; y < board.Height; y++ {
			row[y] = EMPTY
		}
		board.Tiles[x] = row
	}

	// Tiles[x][y] = SNAKE_HEAD or SNAKE_BODY
	for _, snake := range req.Snakes {
		for i, segment := range snake.Coords {
			if !board.Inside(segment.X, segment.Y) {
				continue
			}
			if i == 0 {
				board.Tiles[segment.X][segment.Y] = SNAKE_HEAD
			} else {
				board.Tiles[segment.X][segment.Y] = SNAKE_BODY
			}
		}
	}

	// Tiles[x][y] = FOOD
	for _, food := range req.Food {
		if !board.Inside(food.X, food.Y) {
			continue
		}
		board.Tiles[food.X][food.Y] = FOOD
	}

	return &board
}

func (board *Board) Inside(x, y int) bool {
	return x >= 0 && x < board.Width && y >= 0 && y < board.Height
}

func (board *Board) GetTile(x, y int) Tile {
	if !board.Inside(x, y) {
		return BOUNDARY
	}
	return board.Tiles[x][y]
}

// Represents a 2D coordinate passed as [x, y] in JSON
type Point struct {
	X int
	Y int
}

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
