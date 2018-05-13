package main

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

var mockReq = `{
	"game": {
		"id": "abc"
	},
	"turn": 10,
	"board": {
		"height": 10,
		"width": 10,
		"food": [{
				"x": 1,
				"y": 1
		}],
		"snakes": [{
			"id": "123",
			"name": "snek",
			"health": 1,
			"body": [{
					"x": 1,
					"y": 1
			}]
		}]
	},
	"you": {
		"id": "123",
		"name": "snek",
		"health": 1,
		"body": [{
				"x": 1,
				"y": 1
		}]
	}
}`

func requestWithBody(body string) *http.Request {
	req, err := http.NewRequest("", "", bytes.NewBufferString(body))
	if err != nil {
		panic(err)
	}
	return req
}

func createMockSnakeRequest() *SnakeRequest {
	c := Coord{
		X: 1,
		Y: 1,
	}
	s := Snake{
		ID:     "123",
		Name:   "snek",
		Health: 1,
		Body:   []Coord{c},
	}
	b := Board{
		Height: 10,
		Width:  10,
		Food:   []Coord{c},
		Snakes: []Snake{s},
	}
	g := Game{ID: "abc"}
	sr := SnakeRequest{
		Game:  g,
		Board: b,
		Turn:  10,
		You:   s,
	}
	return &sr
}

func TestDecodeSnakeRequest(t *testing.T) {
	req := requestWithBody(mockReq)
	result, err := DecodeSnakeRequest(req)

	if err != nil {
		t.Fatal(err)
	}
	expected := createMockSnakeRequest()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, result)
	}
}
