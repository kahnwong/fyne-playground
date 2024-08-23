package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func meetings() *fyne.Container {
	// input bindings
	a := binding.NewString()
	err := a.Set("$0")
	if err != nil {
		return nil
	}
	labelA := widget.NewLabelWithData(a)

	// -- output -- //
	output := container.New(layout.NewVBoxLayout(), labelA, widget.NewLabel("$20"), widget.NewLabel("$120"))

	// -- form -- //
	entry := widget.NewEntryWithData(a)
	entry.PlaceHolder = "foo"

	form := container.New(layout.NewVBoxLayout(), entry, widget.NewLabel("fff"), widget.NewLabel("ccc"))

	meetingsSection := container.New(layout.NewHBoxLayout(), output, form)

	return container.New(layout.NewVBoxLayout(), meetingsHeading(), meetingsSection) //, foo)
}

func meetingsHeading() *fyne.Container {
	// constant
	blankLine := canvas.NewText("", color.White)

	// heading
	heading := canvas.NewText("Meeting Cost Calculator", color.White)
	heading.TextSize = 30

	subheading := canvas.NewText("How much does this meeting really cost?", color.White)
	subheading.TextSize = 25

	return container.New(layout.NewVBoxLayout(), heading, blankLine, subheading, blankLine)
}
