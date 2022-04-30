package main

import (
	// "crypto/rand"
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("RandomNumGene")
	w.Resize(fyne.NewSize(600,300))

	label1 := canvas.NewText("Random num gen", color.White)
	label1.TextSize = 40

	btn1 := widget.NewButton("generate",func ()  {
		rand1 := rand.Intn(1000)
		label1.Text = fmt.Sprint(rand1)
		label1.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			btn1,
			label1,
		),
	)

	w.ShowAndRun()
}