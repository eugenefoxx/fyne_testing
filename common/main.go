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
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(400, 320))

	radio := widget.NewRadioGroup([]string{"Radio 1", "Radio 2"}, func(s string) {})

	checks := widget.NewCheckGroup([]string{"Check1", "Check2", "Check3"}, func(s []string) {

	})

	btn2 := widget.NewButton("Жмах", func() {
		//fmt.Println(checks.Selected)
		for _, i := range checks.Selected {
			fmt.Println(i)
		}
		fmt.Println(radio.Selected)
	})

	label := widget.NewLabel("Текст")
	entry := widget.NewEntry()
	btn := widget.NewButton("Жми", func() {
		data := entry.Text
		fmt.Println(data, label.Text)
		label.SetText(data)

	})

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.SetContent(container.NewVBox(
		label,
		entry,
		btn,
		checks,
		btn2,
		radio,
	))
	w.ShowAndRun()
}
