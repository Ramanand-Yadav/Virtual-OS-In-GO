package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main(){

	a:= app.New()
	w := a.NewWindow("line : canvas")
	w.Resize(fyne.NewSize(600,300))

	line1 := canvas.NewLine(color.Black)
	line1.StrokeColor = color.NRGBA{R:255,G:0,B:0,A:255}
	line1.StrokeWidth = 3

	w.SetContent(line1)


	w.ShowAndRun()
}