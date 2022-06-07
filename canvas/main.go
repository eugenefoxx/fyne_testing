package main

import (
	//"fyne.io/fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Подпишись на канал!")
	w.Resize(fyne.NewSize(600, 600))
	img := canvas.NewImageFromFile("predator.jpg")
	w.SetContent(container.NewGridWithColumns(2, img))

	w.ShowAndRun()
}
