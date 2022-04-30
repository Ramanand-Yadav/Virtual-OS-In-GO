package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main(){
	
	a := app.New()
	w := a.NewWindow("Image in fyne")

	//
	w.Resize(fyne.NewSize(600,250))

	//Now Image
	img := canvas.NewImageFromFile("abc.jpg")
	w.SetContent(img)



	w.ShowAndRun()


}