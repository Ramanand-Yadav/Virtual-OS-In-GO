package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main(){

	a := app.New()
	w := a.NewWindow("card")
	w.Resize(fyne.NewSize(600,300))

	//card widgets
	card1 := widget.NewCard(
		"here is my Title",
		"Here is my subtitle",
		// hare we can use any widgets
		canvas.NewImageFromFile("/media/ramanand/UUI/COctober2/fyneTutorial/projects/dice4.png"),
	
	)
	//showwing content
	w.SetContent(card1)

	w.ShowAndRun()
}