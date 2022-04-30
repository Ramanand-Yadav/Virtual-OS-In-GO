package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("NewHSplit")
	w.Resize(fyne.NewSize(600,300))

	//1st widget
	lable1 := canvas.NewText("Text1",color.White)
	lable2 := canvas.NewText("Text2",color.White)
	w1 := widget.NewIcon(theme.CancelIcon())
	w2 := widget.NewButton("allSongs",func(){})
	w.SetContent(
		container.NewHSplit(
			container.NewVBox(
				lable1,
				w1,
			),
			container.NewVBox(
				lable2,
				w2,
			),
		),
	)

	w.ShowAndRun()
}