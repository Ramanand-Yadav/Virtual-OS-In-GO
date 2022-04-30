package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){

	a := app.New()
	w := a.NewWindow("dicee")
	w.Resize(fyne.NewSize(600,300))

	//image 
	img := canvas.NewImageFromFile("dice4.png")
	img.FillMode = canvas.ImageFillOriginal

	//button

	btn1 := widget.NewButton("Play",func ()  {
		rand1 := rand.Intn(6)+1  
		//machine Starts from zero to tackle this problems "use Add One"
		img.File = fmt.Sprintf("dice%d.jpg",rand1)
		img.Refresh()

	})

	//setup content and finish UI
	w.SetContent(
		// newVBox more than one widget
		container.NewVBox(
			img,
			btn1,
		),
		// img,
	)

	w.ShowAndRun()
}