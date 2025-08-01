package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	game.Set()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
