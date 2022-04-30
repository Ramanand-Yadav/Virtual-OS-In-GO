package main

import (
	// "fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

func showBMICalc(){
// func main(){

/*
	a := app.New();
	w := a.NewWindow("BMI Calculator");
	w.Resize(fyne.NewSize(600, 400));
	*/

	labe1 := canvas.NewText("BMI Calculator ", color.Black);
	labe1.Alignment = fyne.TextAlignCenter;
	labe1.TextSize = 20;


	result := canvas.NewText("",color.Black);
	result.Alignment = fyne.TextAlignCenter;
	result.TextSize = 20;

	//input height 
	inputH := widget.NewEntry();
	inputH.SetPlaceHolder("Enter height in cm..");

	//input weight 
	inputW := widget.NewEntry();
	inputW.SetPlaceHolder("Enter weight in KG..");

	btn1 := widget.NewButton("Calc BMI", func ()  {
		h1,_ := strconv.ParseFloat(inputH.Text, 64);
		w,_ := strconv.ParseFloat(inputW.Text, 64);

		result.Text = calculateBMI(h1/100, w);
		result.Refresh();
	});
/*
	w.SetContent(
		container.NewVBox(
			labe1,
			result,
			inputH,
			inputW,
			btn1,
		),
	)

	w.ShowAndRun();
*/

	BMIContainer := container.NewVBox(
			labe1,
			result,
			inputH,
			inputW,
			btn1,
	)
	w:=myApp.NewWindow("Calc");
	w.Resize(fyne.NewSize(600,400));
	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,BMIContainer),
	)
	w.Show();


}
//convetin into func
func calculateBMI(height float64, weight float64) string {
	//BMI calculator Logic 
	//height in cm 
	// height := float64(177.0);//to covert to meter 

	//weight in kg 
	// weight := float64(80.0);

	//BMI formula w/h^2
	var BMI float64 = weight/math.Pow(height,2);
	//math.Pow(base, power)

	//conditions
	//BMI <= 18.4 "You are underweight."
	if(BMI <= 18.4 ){
		// fmt.Println("You are underweight.")
		return "You are underweight.";
	}else if BMI <= 24.9 {
		// fmt.Println("You are healthy.")
		return "You are healthy.";
	}else if BMI <= 29.9 {
		// fmt.Println("You over weight.")
		return "You over weight."
	}else if BMI <= 34.9 {
		// fmt.Println("You are severely over weight.");
		return "You are severely over weight.";
	}else if BMI <= 39.9 {
		// fmt.Println("You are obese.")
		return "You are obese."
	}else{
		// fmt.Println("you are severely obese.")
		return "you are severely obese."
	}
	//BMI <= 24.9 "You are healthy."
	//BMI <= 29.9 "You over weight."
	//BMI <= 34.9 "You are severely over weight."
	//BMI <= 39.9 "You are obese."
	//"you are severely obese."
}





