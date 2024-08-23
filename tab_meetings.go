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

	meetingsSection := container.New(layout.NewHBoxLayout(), output, form)

	return container.New(layout.NewVBoxLayout(), meetingsHeading(), meetingsSection)
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
