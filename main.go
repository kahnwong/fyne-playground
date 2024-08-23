package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Larger")
	w.SetContent(widget.NewLabel("More content"))
	w.Resize(fyne.NewSize(800, 800))
	w.Show()

	a.Run()
}
