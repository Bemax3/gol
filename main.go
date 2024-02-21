package main

import (
	"fmt"

	game "github.com/Bemax3/gol/Game"
)

func main() {
	game := game.NewGame()
	game.FillWithGlider()
	for i := 0; i < 50; i++ {
		game.Display()
		game.NextGeneration()
		fmt.Println("_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _")
	}

}
