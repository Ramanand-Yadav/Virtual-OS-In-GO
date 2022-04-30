package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("Title")

	w.Resize(fyne.NewSize(500, 300))

	// lable1 := widget.NewLabel("FirstName")
	// w.SetContent(lable1)

	//first name button name 
	//second function name 

	btn1 := widget.NewButton("buttonName ",func ()  {
		//button function
		fmt.Println("hello Ramu how are you")
	})

	w.SetContent(btn1)

	w.ShowAndRun()
}