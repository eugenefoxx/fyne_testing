package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	/*btn := container.New(
		layout.NewMaxLayout(),

		canvas.NewRectangle(color.RGBA{112, 40, 144, 1}),
		widget.NewButton("Click me!", func() { fmt.Println("Colored button") }),
	)*/

	w.SetContent(
		//btn,
		container.NewVBox(
			create_colored_button(
				color.RGBA{90, 60, 90, 1},
				widget.NewButton(
					"Click!",
					nil,
				),
			),
			create_colored_button(
				color.RGBA{146, 117, 86, 1},
				widget.NewButton(
					"Lol!",
					nil,
				),
			),
		),
	)
	w.ShowAndRun()
	a.Run()

}

func create_colored_button(c color.RGBA, b *widget.Button) *fyne.Container {
	button := container.New(
		layout.NewMaxLayout(),

		canvas.NewRectangle(c),
		b,
	)

	return button
}
