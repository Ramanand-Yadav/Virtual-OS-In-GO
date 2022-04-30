package main

import (
	// "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
func main(){
	a := app.New();
	w := a.NewWindow("main menu");
	w.Resize(fyne.NewSize(600, 400));

	//first item 
	menuItem := &fyne.Menu{
		Label: "File",
		Items: nil, //we will add some item in next 

	}
	menu := fyne.NewMainMenu(menuItem) //main menu
	//to remove err add &
	//show main menu

	w.SetMainMenu(menu);

	w.ShowAndRun();
}