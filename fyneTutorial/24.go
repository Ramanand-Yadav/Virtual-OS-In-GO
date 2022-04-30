package main

//password generator
//Lower case
//upper case
//Number
//spacial character

import (
	// "fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New();
	w := a.NewWindow("password Generator");
	w.Resize(fyne.NewSize(600, 400));
	//title
	title := canvas.NewText("Password Generator", color.White);
	//input Box
	input := widget.NewEntry();
	input.SetPlaceHolder("Enter password length ");
	//label to show password 
	text := canvas.NewText("", color.White);
	text.TextSize = 20;

	//btn to generate password
	btn1 := widget.NewButton("Generate",func ()  {
		//input
		passlength,_ := strconv.Atoi(input.Text);//convert string to integer
		text.Text = passwordGenerator(passlength);
		text.Refresh();
	});

	w.SetContent(
		container.NewVBox(
			title,
			input,
			text,
			btn1,
		),
	)


	w.ShowAndRun();
}

//converting the code into function
func passwordGenerator(passwordLength int) string{

	//password generator 
	//Lower case
	lowCase := "abcdefghijklmnopqrstuvwxyz";
	//upper case
	upCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ";
	//Number
	Numbers := "0123456789";
	//spacial character
	SpecialChar := "£$&()*+[]@#^-_!?"

	//password length
	// passwordLength := 8;
	//variable for srotin password
	password := "";

	//loop
	for n:=0; n<passwordLength; n++{
		//new randomd character
		rand.Seed(time.Now().UnixNano());
		randNum := rand.Intn(4);
		// fmt.Println(randNum);
		switch(randNum){
			case 0:
				rand.Seed(time.Now().UnixNano());
				randNum := rand.Intn(len(lowCase));
				//len to find length of slice/array
				//Now we will store the generated password character
				password = password+string(lowCase[randNum]);
				//it is byte ..we need string
				//first case completed

			case 1:
				rand.Seed(time.Now().UnixNano());
				randNum := rand.Intn(len(upCase));
				//len to find length of slice/array
				//Now we will store the generated password character
				password = password+string(upCase[randNum]);
				//it is byte ..we need string
				//second case completed

			case 2:
				rand.Seed(time.Now().UnixNano());
				randNum := rand.Intn(len(Numbers));
				//len to find length of slice/array
				//Now we will store the generated password character
				password = password+string(Numbers[randNum]);
				//it is byte ..we need string
				//third case completed
			case 3:
				rand.Seed(time.Now().UnixNano());
				randNum := rand.Intn(len(SpecialChar));
				//len to find length of slice/array
				//Now we will store the generated password character
				password = password+string(SpecialChar[randNum]);
				//it is byte ..we need string
				//4th case completed
		}
	}

	// fmt.Println(password);
	return password;

}
















