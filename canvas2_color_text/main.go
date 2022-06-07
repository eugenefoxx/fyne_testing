package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Подпишись на канал!")
	w.Resize(fyne.NewSize(600, 600))

	color_for := color.NRGBA{R: 23, G: 234, B: 255, A: 255}
	label := canvas.NewText("ЦВЕТ", color_for)

	w.SetContent(label)
	w.ShowAndRun()
}
