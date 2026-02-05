package gooey

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("A Fyne Example")

	myWindow.SetContent(widget.NewLabel("Hello Fyne Gooey Experiments!"))

	myWindow.ShowAndRun()
}
