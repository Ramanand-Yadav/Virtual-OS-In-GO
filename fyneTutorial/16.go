package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a := app.New()
	w := a.NewWindow("IPCityCountry")
	w.Resize(fyne.NewSize(600,300))

	//let create UI First
	label1 := widget.NewLabel("What is my IP")
	label2 := widget.NewLabel("Your IP is ...")
	label_value := widget.NewLabel("...")
	label_City := widget.NewLabel("..")
	label_Country := widget.NewLabel("..")
	btn := widget.NewButton("Run", func ()  {
		//logic
		label_value.Text = myIp()
		label_value.Refresh()
		label_Country.Text = myCountry()
		label_Country.Refresh()
		label_City.Text = myCity()
		label_City.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			label1,
			label2,
			label_value,
			label_Country,
			label_City,
			btn,
		),
	)

	w.ShowAndRun()
}

func myIp() string{
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil{
		return err.Error()
	}

	defer req.Body.Close()

	body,err := ioutil.ReadAll(req.Body)
	if err!= nil{
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body,&ip)

	return ip.Query
}

func myCity() string{
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil{
		return err.Error()
	}

	defer req.Body.Close()

	body,err := ioutil.ReadAll(req.Body)
	if err!= nil{
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body,&ip)

	return ip.City
}

func myCountry() string{
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil{
		return err.Error()
	}

	defer req.Body.Close()

	body,err := ioutil.ReadAll(req.Body)
	if err!= nil{
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body,&ip)

	return ip.Country
}

type IP struct{
	Query string
	Country string
	City string
}