package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Programm!")
	w.Resize(fyne.NewSize(500, 500))

	item1 := fyne.NewMenuItem("Actions", nil)
	item2 := fyne.NewMenuItem("Say", nil)

	item1.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Print", func() { fmt.Println("test Print") }),
		fyne.NewMenuItem("Save", func() { fmt.Println("test Save") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("test Copy") }),
	)

	item2.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Hi", func() { fmt.Println("test Hi") }),
		fyne.NewMenuItem("Bye", func() { fmt.Println("test Bye") }),
		fyne.NewMenuItem("Lol", func() { fmt.Println("test Lol") }),
	)
	menu := fyne.NewMenu("Buttons", item1, item2)

	main := fyne.NewMainMenu(menu)

	w.SetMainMenu(main)
	w.ShowAndRun()
	a.Run()

}
