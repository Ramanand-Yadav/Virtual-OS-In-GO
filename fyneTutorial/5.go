package main

import (
	"net/url"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main(){

	a := app.New()

	w := a.NewWindow("Title")
	//creating hyperLink

	url, _ := url.Parse("https://google.com")
	//first value is lable 
	//2nd value link 
	HyperLink1 := widget.NewHyperlink("Visit me",url)
	w.SetContent(HyperLink1)


	w.ShowAndRun()

}