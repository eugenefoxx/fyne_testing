package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	card := widget.NewCard(
		"IT Еж Group",
		"Ознакмтесь с пользовательским солашением",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Параграф 1",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Пункт 1",
						widget.NewLabel("Информация о пункте 1"),
					),
					widget.NewAccordionItem(
						"Пункт 2",
						widget.NewLabel("Информация о пункте 2"),
					),
					widget.NewAccordionItem(
						"Пункт 3",
						widget.NewLabel("Информация о пункте 3"),
					),
				),
			),
			widget.NewAccordionItem(
				"Параграф 1",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Пункт 1",
						widget.NewLabel("Информация о пункте 1"),
					),
					widget.NewAccordionItem(
						"Пункт 2",
						widget.NewLabel("Информация о пункте 2"),
					),
					widget.NewAccordionItem(
						"Пункт 3",
						widget.NewLabel("Информация о пункте 3"),
					),
				),
			),
			widget.NewAccordionItem(
				"Дополнтительная информация",
				widget.NewLabel("Доп инфомация дляпользователя"),
			),
		),
	)

	w.SetContent(card)
	w.ShowAndRun()
	a.Run()
}
