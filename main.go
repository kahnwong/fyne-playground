package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// init
	a := app.New()
	w := a.NewWindow("Fyne Playground")

	// init tabs
	tabs := container.NewAppTabs()

	// tab: meetings
	tabs.Append(container.NewTabItem("Meetings", widget.NewLabel("Hello")))

	// set tab location
	tabs.SetTabLocation(container.TabLocationLeading)

	// set window
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1024, 768))

	// run app
	w.ShowAndRun()
}
