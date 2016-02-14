package main

import (
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestNewBoard(t *testing.T) {
	requests := []MoveRequest{
		{
			Snakes: []Snake{},
			Food:   []Point{},
			Width:  20,
			Height: 10,
		},
		{
			Snakes: []Snake{
				{ID: "1", Coords: []Point{
					{19, 9}, {18, 9}, {17, 9},
				}},
				{ID: "2", Coords: []Point{
					{0, 0}, {0, 1}, {0, 2},
				}},
			},
			Food:   []Point{{6, 4}, {8, 3}},
			Width:  20,
			Height: 10,
		},
	}

	for _, request := range requests {
		board := NewBoard(&request)

		assert.Equal(t, board.Width, request.Width)
		assert.Equal(t, board.Height, request.Height)

		assert.Equal(t, len(board.Tiles), board.Width)
		for _, row := range board.Tiles {
			assert.Equal(t, len(row), board.Height)
		}

		for _, snake := range request.Snakes {
			head := snake.Coords[0]
			tile := board.Tiles[head.X][head.Y]
			assert.Equal(t, tile, SNAKE_HEAD)
		}
		for _, food := range request.Food {
			tile := board.Tiles[food.X][food.Y]
			assert.Equal(t, tile, FOOD)
		}
	}
}

func TestBoardInside(t *testing.T) {
	board := Board{
		Width:  10,
		Height: 20,
	}

	assert.Equal(t, board.Inside(0, 0), true)
	assert.Equal(t, board.Inside(9, 0), true)
	assert.Equal(t, board.Inside(0, 19), true)

	assert.Equal(t, board.Inside(-1, 0), false)
	assert.Equal(t, board.Inside(0, -1), false)
	assert.Equal(t, board.Inside(10, 0), false)
	assert.Equal(t, board.Inside(0, 20), false)
}
