package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	text := "Occaecat esse dolor eiusmod voluptate cillum sunt sunt qui nisi ipsum aliquip nulla excepteur. Mollit id culpa aliqua velit nostrud est sit adipisicing voluptate quis fugiat ipsum nulla sint. Lorem enim consectetur elit deserunt qui in et tempor aliquip ipsum occaecat aliquip. Commodo do mollit esse aliquip culpa amet et in duis nostrud adipisicing."
	text_label := widget.NewLabel(text)
	text_label.Wrapping = fyne.TextWrapBreak
	item1 := widget.NewAccordionItem(
		"Read more info",
		//widget.NewLabel("Info"),
		text_label,
	)

	item2 := widget.NewAccordionItem(
		"Button",
		widget.NewButton(
			"Say Hello!",
			func() {
				fmt.Println("Hello")
			},
		),
	)

	item3 := widget.NewAccordionItem(
		"Chapters",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Chapter 1",
				widget.NewLabel("Some text here ..."),
			),
			widget.NewAccordionItem(
				"Chapter 2",
				widget.NewLabel("Some2 text here ..."),
			),
			widget.NewAccordionItem(
				"Chapter 3",
				widget.NewLabel("Some3 text here ..."),
			),
		),
	)

	accordion := widget.NewAccordion(item1, item2, item3)

	w.SetContent(accordion)
	w.ShowAndRun()
	a.Run()
}
