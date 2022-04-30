package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("Title")

	w.Resize(fyne.NewSize(500, 100))

	// leble1 := widget.NewLabel("Enter Your Name")
	// w.SetContent(leble1)

	checkBox1 := widget.NewCheck("Male", func(b bool) {
		// fmt.Println(fmt.Sprintf("my check box valu is &t",b))
		if b==true{
			fmt.Println("male")
		}else{
			fmt.Println("Not male")
		}
	})

	w.SetContent(checkBox1)


	w.ShowAndRun()
}