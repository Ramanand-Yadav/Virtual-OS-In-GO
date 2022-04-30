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
	w := a.NewWindow("theme")
	//set theme
	
	
	w.Resize(fyne.NewSize(600, 300))

	label1 := canvas.NewText("fyne themes",color.White)
	label1.TextSize = 50
	lbtn := widget.NewButton("Light",func ()  {
		label1.Color = color.Black
		a.Settings().SetTheme(theme.LightTheme())
	})
	dbtn := widget.NewButton("Dark",func ()  {
		label1.Color = color.White
		a.Settings().SetTheme(theme.DarkTheme())
	})
	ebtn := widget.NewButton("exits",func ()  {
		a.Quit()
	})

	w.SetContent(
		container.NewVBox(
			label1,
			lbtn,
			dbtn,
			ebtn,
		),
	)



	w.ShowAndRun()
}