package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main(){

	a := app.New()

	w := a.NewWindow("Title")
	w.Resize(fyne.NewSize(500, 200))
	
	//Text widgets in Canvas
	//first value for label / action
	//2nd value for color
	//Now creating color
	
	colorx := color.NRGBA{R:0, G:0, B:0, A:255}
	// R is Red (0 to 255) 
	//G is Green ( 0 to 255 )
	//A is for alpha for capacity 
	//b is blue
	textX := canvas.NewText("Here is my text", colorx)
	w.SetContent(textX)

	w.ShowAndRun()

}