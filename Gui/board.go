package gui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	game "github.com/Bemax3/gol/Game"
	settings "github.com/Bemax3/gol/Settings"
)

type Board struct {
	Container *fyne.Container
	Game      *game.Game
	Cells     [][]*canvas.Rectangle
}

func NewBoard(g *game.Game, c *fyne.Container) *Board {
	b := &Board{
		Container: c,
		Game:      g,
		Cells:     make([][]*canvas.Rectangle, settings.Dimension),
	}

	for i := 0; i < settings.Dimension; i++ {
		b.Cells[i] = make([]*canvas.Rectangle, settings.Dimension)
		for j := 0; j < settings.Dimension; j++ {
			rect := canvas.NewRectangle(color.Black)
			rect.Move(fyne.NewPos(float32(i*settings.UnitSize), float32(j*settings.UnitSize)))
			rect.Resize(fyne.NewSize(settings.UnitSize-1, settings.UnitSize-1))
			b.Cells[i][j] = rect
			c.Add(rect)
		}
	}
	return b
}

func (b *Board) DrawLines() {
	var line *canvas.Line
	for i := 0; i <= settings.Dimension; i++ {
		line = canvas.NewLine(color.White)
		line.Position1 = fyne.NewPos(float32(i*settings.UnitSize), 0)
		line.Position2 = fyne.NewPos(float32(i*settings.UnitSize), settings.BoardHeight)
		b.Container.Add(line)
		line = canvas.NewLine(color.White)
		line.Position1 = fyne.NewPos(0, float32(i*settings.UnitSize))
		line.Position2 = fyne.NewPos(settings.BoardWidth, float32(i*settings.UnitSize))
		b.Container.Add(line)
	}
}

func (b *Board) Play() {
	go func() {
		for b.Game.Running {
			b.Game.NextGeneration()
			b.Repaint()
			time.Sleep(time.Duration(b.Game.Delay) * time.Millisecond)
		}
	}()
}

func (b *Board) Start() {
	if b.Game.Running {
		return
	}
	b.Game.Running = true
	b.Play()
	b.Repaint()
}

func (b *Board) Stop() {
	if !b.Game.Running {
		return
	}
	b.Game.Running = false
	b.Repaint()
}

func (b *Board) SpeedUp() {

	b.Game.Delay -= 10
	if b.Game.Delay < 0 {
		b.Game.Delay = 0
	}
	b.Repaint()

}

func (b *Board) SlowDown() {
	b.Game.Delay += 10
	b.Repaint()
}

func (b *Board) Reset() {
	// for i := 0; i < settings.Dimension; i++ {
	// 	for j := 0; j < settings.Dimension; j++ {
	// 		b.Game.Cells[i][j] = 0
	// 	}
	// }
	b.Game.FillRandomly()
	b.Repaint()
}

func (b *Board) Repaint() {
	for i := 0; i < settings.Dimension; i++ {
		for j := 0; j < settings.Dimension; j++ {
			if b.Game.Cells[i][j] == 1 {
				b.Cells[i][j].FillColor = color.White
			} else {
				b.Cells[i][j].FillColor = color.Black
			}
			b.Cells[i][j].Refresh()
		}
	}
	b.Container.Refresh()
}
