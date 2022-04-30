package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
func main(){
	a := app.NewWithID("com.anydom.com");
	w := a.NewWindow("menu & submenu");
	w.Resize(fyne.NewSize(600, 400));
	menuItem1 := fyne.NewMenuItem("New",func() {fmt.Println("new Pressed")});
	menuItem2 := fyne.NewMenuItem("Edit",func() {fmt.Println("edit Pressed")});
	menuItem3 := fyne.NewMenuItem("Save",func() {fmt.Println("save Pressed")});
	menuItem4 := fyne.NewMenuItem("Run",func() {fmt.Println("run pressed")});
	menuItem5 := fyne.NewMenuItem("Help",func() {fmt.Println("hep Pressed")});

	newMenu1 := fyne.NewMenu("File", menuItem1,menuItem2,menuItem3);
	newMenu2 := fyne.NewMenu("Other", menuItem1,menuItem4,menuItem2);
	newMenu3 := fyne.NewMenu("Help", menuItem5);

	menu := fyne.NewMainMenu(newMenu1,newMenu2,newMenu3);
	w.SetMainMenu(menu);



	w.ShowAndRun();
}