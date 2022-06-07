package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	// VScroll
	// HScroll

	scr := container.NewVScroll(
		container.NewVBox(
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
		),
	)

	scr.SetMinSize(fyne.NewSize(200, 150))

	w.SetContent(
		container.NewVBox(
			scr,
			widget.NewLabel("TEXT HERE"),
		),
	)
	w.ShowAndRun()
	a.Run()

}
