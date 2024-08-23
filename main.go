package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func meetings() *fyne.Container {
	// constant
	blankLine := canvas.NewText("", color.White)

	// header
	heading := canvas.NewText("Meeting Cost Calculator", color.White)
	heading.TextSize = 30

	subheading := canvas.NewText("How much does this meeting really cost?", color.White)
	subheading.TextSize = 25

	content := container.New(layout.NewVBoxLayout(), heading, blankLine, subheading, blankLine)

	// meetingsSection
	// -- output -- //
	output := container.New(layout.NewVBoxLayout(), widget.NewLabel("$200"), widget.NewLabel("$20"), widget.NewLabel("$120"))

	// -- form -- //
	str := binding.NewString()
	err := str.Set("Hi!")
	if err != nil {
		return nil
	}

	form := container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	)

	meetingsSection := container.New(layout.NewHBoxLayout(), output, blankLine, form)

	return container.New(layout.NewVBoxLayout(), content, meetingsSection)
}

func main() {
	// init
	a := app.New()
	w := a.NewWindow("Fyne Playground")

	// init tabs
	tabs := container.NewAppTabs()

	// tab: meetings
	tabs.Append(container.NewTabItem("Meetings", meetings()))

	// set tab location
	tabs.SetTabLocation(container.TabLocationLeading)

	// set window
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1024, 768))

	// run app
	w.ShowAndRun()
}
