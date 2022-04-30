package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


func main(){

	a := app.New()
	w := a.NewWindow("Title")

	w.Resize(fyne.NewSize(600, 400))

	w.ShowAndRun()
}