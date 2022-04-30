package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)


func main(){

	a:= app.New()
	w:= a.NewWindow("Draw Circle")
	w.Resize(fyne.NewSize(600, 300))
	

	//

	// circleX := canvas.NewCircle(color.Black)
	circleX := canvas.NewCircle(color.NRGBA{R:0,G:0, B:255,A:255 })
	circleX.StrokeColor = color.NRGBA{R:255,G:0,B:0,A:255}
	circleX.StrokeWidth= 4

	w.SetContent(circleX)


	w.ShowAndRun()
}