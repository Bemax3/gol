package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	MenuWidth  = 200
	MenuHeight = 800
)

func Menu(b *Board) *fyne.Container {
	startBtn := widget.NewButton("Start", b.Start)
	stopBtn := widget.NewButton("Stop", b.Stop)
	speedBtn := widget.NewButton("Speed Up", b.SpeedUp)
	slowBtn := widget.NewButton("Slow Down", b.SlowDown)
	resetBtn := widget.NewButton("Reset", b.Reset)
	c := container.NewVBox(
		startBtn, stopBtn, speedBtn, slowBtn, resetBtn,
	)
	c.Resize(fyne.NewSize(MenuWidth, MenuHeight))
	return c
}
