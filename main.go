package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	worldTurnNumber int
}

type Human struct {
	id      string
	name    string
	satiety int
}

func runMainLoop(game *Game) {
	for {
		interval := time.Microsecond * 16666
		time.Sleep(interval)

		game.worldTurnNumber = game.worldTurnNumber + 1

		fmt.Println(game.worldTurnNumber)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}

	go runMainLoop(game)

	fmt.Println("Start a server.")

	time.Sleep(time.Second * 5)
}
