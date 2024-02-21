package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	game "github.com/Bemax3/gol/Game"
	gui "github.com/Bemax3/gol/Gui"
	settings "github.com/Bemax3/gol/Settings"
)

func main() {
	a := app.New()
	w := a.NewWindow("Game of Life")
	w.Resize(fyne.NewSize(settings.BoardWidth+gui.MenuWidth, settings.BoardHeight+gui.MenuHeight))

	game := game.NewGame(100)
	game.FillWithGliderGun()

	c := container.NewWithoutLayout()
	c.Resize(fyne.NewSize(settings.BoardWidth, settings.BoardHeight))

	b := gui.NewBoard(game, c)

	w.SetContent(container.NewHBox(
		gui.Menu(b),
		b.Container,
	))
	b.Play()
	w.ShowAndRun()
}
