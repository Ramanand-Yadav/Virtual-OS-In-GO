package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main(){

	a := app.New();
	w := a.NewWindow("Loading image from internet");
	w.Resize(fyne.NewSize(600, 400));

	r,_ := fyne.LoadResourceFromURLString("https://picsum.photos/200");
	img := canvas.NewImageFromResource(r);

	w.SetContent(img);
	w.ShowAndRun();
}