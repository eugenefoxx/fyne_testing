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

	title := widget.NewLabel("Орфография")

	right_label := widget.NewLabel("Правильно")
	incorrect_label := widget.NewLabel("Не правильно")

	r_text := widget.NewLabel("Magna culpa nostrud sit excepteur. Id exercitation ullamco reprehenderit ex qui in sit. Enim nisi mollit aute id et culpa. Non cillum adipisicing sunt do ut duis veniam.")
	r_text.Wrapping = fyne.TextWrapBreak

	i_text := widget.NewLabel("Magna culpa nostrud sit excepteur. Id exercitation ullamco reprehenderit ex qui in sit. Enim nisi mollit aute id et culpa. Non cillum adipisicing sunt do ut duis veniam.")
	i_text.Wrapping = fyne.TextWrapBreak

	r_box := container.NewVBox(right_label, r_text)
	i_box := container.NewVBox(incorrect_label, i_text)

	split := container.NewHSplit(r_box, i_box)

	cont := container.NewVBox(title, split)

	w.SetContent(cont)
	w.ShowAndRun()
	a.Run()
}
