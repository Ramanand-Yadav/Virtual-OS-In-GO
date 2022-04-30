package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){

	a := app.New()

	w := a.NewWindow("Layout")
	w.Resize(fyne.NewSize(600, 400))
	//new btn
	btn1 := widget.NewButton("click me",func ()  {
		
	})
	label1 := widget.NewLabel("here is my text")

	//newVbox
	box1 := container.NewVBox(
		label1,
		btn1,
	)
	//NewHBox
	// box2 := container.NewHBox(
	// 	btn1,
	// 	label1,
	// )
	w.SetContent(
		box1,
		// box2,
	)

	w.ShowAndRun()
}