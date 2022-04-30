package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
func main(){
	a := app.NewWithID("com.anydom.com");
	w := a.NewWindow("Menu With child menu");
	w.Resize(fyne.NewSize(600,400));

	menuItem1 := fyne.NewMenuItem("Edit",func() {});
	menuItem2 := fyne.NewMenuItem("Details",func() {});
	menuItem3 := fyne.NewMenuItem("Home",func() {});
	menuItem4 := fyne.NewMenuItem("Run",func() {});

	menuItem1.ChildMenu = fyne.NewMenu(
		"",//leave label empty
		fyne.NewMenuItem("copy",nil),
		fyne.NewMenuItem("cut",nil),
		fyne.NewMenuItem("paste",nil),
	)
	menuItem2.ChildMenu = fyne.NewMenu(
		"",//leave label empty
		fyne.NewMenuItem("book",nil),
		fyne.NewMenuItem("magzine",nil),
		fyne.NewMenuItem("notebook",nil),
	)
	menuItem3.ChildMenu = fyne.NewMenu(
		"",//leave label empty
		fyne.NewMenuItem("school",nil),
		fyne.NewMenuItem("college",nil),
		fyne.NewMenuItem("university",nil),
	)
	

	newMenu1 := fyne.NewMenu("File", menuItem1,menuItem2,menuItem3,menuItem4);
	newMenu2 := fyne.NewMenu("Help",menuItem1,menuItem2,menuItem3,menuItem4);

	menu := fyne.NewMainMenu(newMenu1,newMenu2);
	w.SetMainMenu(menu);



	w.ShowAndRun();
}