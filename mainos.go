package main

// for running the whole code all .go in one folder and 
//  and then run the command -> (go run ./) or (go run ./mainos.go ........allfiles name)

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("pepos");
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var img fyne.CanvasObject;
var DeskBtn fyne.Widget
var panelContent *fyne.Container

func main(){
	myApp.Settings().SetTheme(theme.LightTheme());
	img = canvas.NewImageFromFile("/media/ramanand/UUI/COctober2/fyneProject/osimage.jpg");
	btn1 = widget.NewButtonWithIcon("Weather",theme.InfoIcon(),func ()  {
		showWeatherApp(myWindow);
	})
	btn2 = widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(),func ()  {
		showCalc();
	})
	btn3 = widget.NewButtonWithIcon("Gallery",theme.StorageIcon(),func ()  {
		showGalleryApp(myWindow);
	})
	btn4 = widget.NewButtonWithIcon("Text Editor",theme.DocumentIcon(),func ()  {
		showTextEditor();
	})

	btn5 = widget.NewButtonWithIcon("BMI",theme.ConfirmIcon(),func ()  {
		showBMICalc();
	})

	btn6 = widget.NewButtonWithIcon("NotePad",theme.DocumentCreateIcon(),func ()  {
		showNotePadApp();
	})
	
	DeskBtn = widget.NewButtonWithIcon("Home",theme.HomeIcon(),func ()  {
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(
		6,
		DeskBtn,
		btn1,
		btn2,
		btn3,
		btn4,
		btn5,
		btn6,
	))

	myWindow.Resize(fyne.NewSize(1280,720));
	myWindow.CenterOnScreen();
	myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img),)

	myWindow.ShowAndRun();



}





