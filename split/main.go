package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	// HSplit
	// VSplit
	label1 := widget.NewLabel("Action 1")
	btn1 := widget.NewButton("Do 1!", func() {
		fmt.Println("1")
	})
	label2 := widget.NewLabel("Action 2")
	btn2 := widget.NewButton("Do 2!", func() {
		fmt.Println("2")
	})

	w.SetContent(
		container.NewHSplit(
			//label1,
			//label2,
			container.NewVBox(
				label1,
				btn1,
			),
			container.NewVBox(
				label2,
				btn2,
			),
		),
	)
	w.ShowAndRun()
	a.Run()
}
