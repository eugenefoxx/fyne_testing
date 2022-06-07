package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("ПОДПИШИСЬ НА КАНАЛ!")
	w.Resize(fyne.NewSize(500, 500))

	// Главное меню - дочернее меню - кнопки
	file_item1 := fyne.NewMenuItem("New File", func() {
		os.Create("created.txx")
	})
	file_item2 := fyne.NewMenuItem("Save", func() {
		fmt.Println("File saved!")
	})
	menu1 := fyne.NewMenu("File", file_item1, file_item2)

	actions_item1 := fyne.NewMenuItem("Hello!", func() {
		fmt.Println("Hello from menu!")
	})
	actions_item2 := fyne.NewMenuItem("Bye!", func() {
		fmt.Println("Bye from menu!")
	})
	actions_item3 := fyne.NewMenuItem("Button!", func() {
		fmt.Println("Clicked!")
	})
	menu2 := fyne.NewMenu("Actions", actions_item1, actions_item2, actions_item3)

	main_menu := fyne.NewMainMenu(menu1, menu2)
	w.SetMainMenu(main_menu)

	w.ShowAndRun()

}
