package main

import (
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
	w := a.NewWindow("Random Color ")
	w.Resize(fyne.NewSize(400,400))
	colorx := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	rect1 := canvas.NewRectangle(colorx)
	rect1.SetMinSize(fyne.NewSize(400,200))
	btn1 := widget.NewButton("Random color",func ()  {
		//unit8 is nacessary to convert int to uint8
		rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
		rect1.Refresh()
	})
	btn2 := widget.NewButton("Red Random color",func ()  {
		rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: 0, B: 0, A: 255}
		rect1.Refresh()
	})
	btn3 := widget.NewButton("Green Random color",func ()  {
		rect1.FillColor = color.NRGBA{R: 0, G: uint8(rand.Intn(255)), B: 0, A: 255}
		rect1.Refresh()
	})
	btn4 := widget.NewButton("Blue Random color",func ()  {
		rect1.FillColor = color.NRGBA{R: 0, G: 0, B: uint8(rand.Intn(255)), A: 255}
		rect1.Refresh()
	})
	btn5 := widget.NewButton("exits",func() {
		a.Quit()
	})
	w.SetContent(
		container.NewVBox(
			rect1,
			btn1,
			btn2,
			btn3,
			btn4,
			btn5,
		),
	)

	w.ShowAndRun()
}