package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main(){
	
	a:= app.New()

	w:= a.NewWindow("icon ")
	w.Resize(fyne.NewSize(600,300))
	icon := widget.NewIcon(theme.CancelIcon())

	w.SetContent(icon)


	w.ShowAndRun()
}