package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
func main(){
	a := app.New();
	w := a.NewWindow("menu ahead");
	w.Resize(fyne.NewSize(600,400));

	menuItem1 := fyne.NewMenuItem("New", func ()  {
		fmt.Println("New");
	});
	menuItem2 := fyne.NewMenuItem("Save", func ()  {
		fmt.Println("Save");
	});
	menuItem3 := fyne.NewMenuItem("edit",nil);

	//New menu
	newMenu := fyne.NewMenu("file",menuItem1,menuItem2,menuItem3);
	menu := fyne.NewMainMenu(newMenu);
	w.SetMainMenu(menu);

	w.ShowAndRun();
}