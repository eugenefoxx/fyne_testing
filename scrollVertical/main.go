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

	title := widget.NewLabel("CHOOSE THE OPTION")

	ch := widget.NewCheckGroup(
		[]string{
			"Option 1",
			"Option 2",
			"Option 3",
			"Option 4",
			"Option 5",
			"Option 6",
			"Option 7",
			"Option 8",
			"Option 9",
			"Option 10",
		},
		nil,
	)

	scr := container.NewVScroll(
		container.NewVBox(
			ch,
		),
	)
	scr.SetMinSize(fyne.NewSize(400, 150))

	res := widget.NewLabel("")

	w.SetContent(
		container.NewVBox(
			title,
			scr,
			widget.NewButton(
				"Submit",
				func() {
					res.SetText("")

					for _, i := range ch.Selected {
						res.SetText(res.Text + i + "\n")
					}
				},
			),
			res,
		),
	)

	w.ShowAndRun()
	a.Run()

}
