package game

import (
	"math/rand"
)

type Game struct {
	board [20][20]int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) FillRandomly() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			g.board[i][j] = rand.Intn(2)
		}
	}
}
func (g *Game) FillWithGlider() {
	g.board[0][1] = 1
	g.board[1][2] = 1
	g.board[2][0] = 1
	g.board[2][1] = 1
	g.board[2][2] = 1
}

func (g *Game) Display() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if g.board[i][j] == 0 {
				print(" ")
			} else {
				print("X")
			}
			print(" ")
		}
		println()
	}
}

func (g *Game) countAliveNeighbors(x, y int) int {
	alive := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i > 19 || y+j < 0 || y+j > 19 {
				continue
			}
			alive += g.board[x+i][y+j]
		}
	}
	return alive
}

func (g *Game) NextGeneration() {
	next := [20][20]int{}
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			alive := g.board[i][j]
			neighbors := g.countAliveNeighbors(i, j)
			if alive == 1 && (neighbors < 2 || neighbors > 3) {
				next[i][j] = 0
			} else if alive == 0 && neighbors == 3 {
				next[i][j] = 1
			} else {
				next[i][j] = alive
			}
		}
	}
	g.board = next
}
