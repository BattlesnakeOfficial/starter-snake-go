package main

import (
	"log"
	"net/http"
)

func Start(res http.ResponseWriter, req *http.Request) {
	data, err := DecodeSnakeRequest(req)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(data)

	respond(res, StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	data, err := DecodeSnakeRequest(req)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(data)

	respond(res, MoveResponse{
		Move: "down",
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}
