package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main(){

	a := app.New()
	w := a.NewWindow("Title")

	w.Resize(fyne.NewSize(500, 200))

	// w.SetContent(widget.NewLabel("FirstName"))

	label2 := widget.NewLabel("LastName")
	w.SetContent(label2)


	w.ShowAndRun()
}