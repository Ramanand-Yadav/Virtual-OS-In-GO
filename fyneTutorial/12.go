package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main(){

	a := app.New()

	w := a.NewWindow("gradient : video")

	w.Resize(fyne.NewSize(600, 300))

	//Gradient
	// gradient1 := canvas.NewHorizontalGradient(color.Black,color.White)
	// gradient2 := canvas.NewVerticalGradient(color.Black, color.White)
	// gradient3 := canvas.NewLinearGradient(color.Black,color.White,45)  45 degree angle

	gradient4 := canvas.NewRadialGradient(color.White,color.NRGBA{R:255,G:0,B:0,A:255})

	w.SetContent(gradient4)

	w.ShowAndRun()
}