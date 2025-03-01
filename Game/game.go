package game

import (
	"math/rand"
	"sync"

	s "github.com/Bemax3/gol/Settings"
)

type Game struct {
	Running   bool
	Delay     int
	Cells     [][]int
	tempCells [][]int
}

func NewGame(Delay int) *Game {
	cells := make([][]int, s.Dimension)
	tempCells := make([][]int, s.Dimension)
	for i := 0; i < s.Dimension; i++ {
		cells[i] = make([]int, s.Dimension)
		tempCells[i] = make([]int, s.Dimension)
	}
	return &Game{
		Running:   true,
		Delay:     Delay,
		Cells:     cells,
		tempCells: tempCells,
	}
}

func (g *Game) FillRandomly() {
	for i := 0; i < s.Dimension; i++ {
		for j := 0; j < s.Dimension; j++ {
			g.Cells[i][j] = rand.Intn(2)
		}
	}
}

func (g *Game) FillWithGliderGun() {
	g.Cells[1][5] = 1
	g.Cells[1][6] = 1
	g.Cells[2][5] = 1
	g.Cells[2][6] = 1
	g.Cells[11][5] = 1
	g.Cells[11][6] = 1
	g.Cells[11][7] = 1
	g.Cells[12][4] = 1
	g.Cells[12][8] = 1
	g.Cells[13][3] = 1
	g.Cells[13][9] = 1
	g.Cells[14][3] = 1
	g.Cells[14][9] = 1
	g.Cells[15][6] = 1
	g.Cells[16][4] = 1
	g.Cells[16][8] = 1
	g.Cells[17][5] = 1
	g.Cells[17][6] = 1
	g.Cells[17][7] = 1
	g.Cells[18][6] = 1
	g.Cells[21][3] = 1
	g.Cells[21][4] = 1
	g.Cells[21][5] = 1
	g.Cells[22][3] = 1
	g.Cells[22][4] = 1
	g.Cells[22][5] = 1
	g.Cells[23][2] = 1
	g.Cells[23][6] = 1
	g.Cells[25][1] = 1
	g.Cells[25][2] = 1
	g.Cells[25][6] = 1
	g.Cells[25][7] = 1
	g.Cells[35][3] = 1
	g.Cells[35][4] = 1
	g.Cells[36][3] = 1
	g.Cells[36][4] = 1
}

func (g *Game) FillWithGlider() {
	g.Cells[0][1] = 1
	g.Cells[1][2] = 1
	g.Cells[2][0] = 1
	g.Cells[2][1] = 1
	g.Cells[2][2] = 1
}

func (g *Game) countAliveNeighbors(x, y int) int {
	alive := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= s.Dimension || y+j < 0 || y+j >= s.Dimension {
				continue
			}
			alive += g.Cells[x+i][y+j]
		}
	}
	return alive
}

func (g *Game) NextGeneration() {
	var wg sync.WaitGroup
	wg.Add(s.Dimension)

	for i := 0; i < s.Dimension; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < s.Dimension; j++ {
				neighbors := g.countAliveNeighbors(i, j)
				if g.Cells[i][j] == 1 {
					if neighbors < 2 || neighbors > 3 {
						g.tempCells[i][j] = 0
					} else {
						g.tempCells[i][j] = 1
					}
				} else {
					if neighbors == 3 {
						g.tempCells[i][j] = 1
					} else {
						g.tempCells[i][j] = 0
					}
				}
			}
		}(i)
	}

	wg.Wait()

	g.Cells, g.tempCells = g.tempCells, g.Cells
}
