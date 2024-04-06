package main

import game "github.com/damienjacinto/connect4/internal"

func main() {
	game := game.NewGame(640, 480, "Connect 4")
	game.Start()
}
