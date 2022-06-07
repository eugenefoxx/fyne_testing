package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	rec1 := canvas.NewRectangle(color.RGBA{79, 32, 106, 1})
	rec2 := canvas.NewRectangle(color.RGBA{157, 51, 51, 1})
	rec3 := canvas.NewRectangle(color.RGBA{118, 183, 39, 1})

	rec1.SetMinSize(fyne.NewSize(400, 150))
	rec2.SetMinSize(fyne.NewSize(400, 150))
	rec3.SetMinSize(fyne.NewSize(400, 150))

	scr := container.NewHScroll(
		container.NewHBox(
			rec1,
			rec2,
			rec3,
		),
	)

	w.SetContent(scr)
	w.ShowAndRun()
	a.Run()

}
