package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Start(res http.ResponseWriter, req *http.Request) {
	log.Print("START REQUEST")

	data, err := NewStartRequest(req)
	if err != nil {
		log.Printf("Bad start request: %v", err)
		respond(res, StartResponse{})
	}
	dump(data)

	respond(res, StartResponse{
		Taunt: "battlesnake-go!",
		Color: "#00FF00",
		Name:  "battlesnake-go",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	log.Printf("MOVE REQUEST")

	data, err := NewMoveRequest(req)

	if err != nil {
		log.Printf("Bad move request: %v", err)
		respond(res, MoveResponse{
			Move: "up",
		})
		return
	}
	dump(data)

	directions := []string{
		"up",
		"down",
		"left",
		"right",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	respond(res, MoveResponse{
		Move: directions[r.Intn(4)],
	})
}

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}

func dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}
