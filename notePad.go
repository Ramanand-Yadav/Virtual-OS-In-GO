package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)
var count1 int = 1
var filepath string
//change main -> showNotePadApp
func showNotePadApp(){
	/*
	a := app.New()
	w := a.NewWindow("NotePad")
	w.Resize(fyne.NewSize(600,400))
	*/

	w:=myApp.NewWindow("NotePad");
	w.Resize(fyne.NewSize(600,400));

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text ...")
	input.Resize(fyne.NewSize(600,400))

	new1 := fyne.NewMenuItem("New",func() {
		filepath = ""
		w.SetTitle("Fyne Note Pad")
		input.Text = ""
		input.Refresh()
	})

	save1 := fyne.NewMenuItem("Save",func() {
		if filepath != "" {
			f , err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
			if err != nil {

			}
			defer f.Close()
			f.WriteString(input.Text)
		}else{
			saveFileDialog := dialog.NewFileSave(
				func (r fyne.URIWriteCloser, _ error)  {
					textData := []byte(input.Text)
					r.Write(textData)
					filepath = r.URI().Path()
					w.SetTitle(filepath)
				},w)
			saveFileDialog.SetFileName("NewFile"+strconv.Itoa(count1-1)+".text")
			count1 = count1 + 1
			saveFileDialog.Show()
		}
	})

	saveAs1 := fyne.NewMenuItem("Save as..",func() {
		saveFileDialog := dialog.NewFileSave(
			func (r fyne.URIWriteCloser, _ error)  {
				textData := []byte(input.Text)
				r.Write(textData)
				filepath = r.URI().Path()
				w.SetTitle(filepath)
			},w)
		saveFileDialog.SetFileName("NewFile"+strconv.Itoa(count1-1)+".text")
		count1 = count1 + 1
		saveFileDialog.Show()
	})

	open1 := fyne.NewMenuItem("Open",func() {
		openFileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(uc)
				result := fyne.NewStaticResource("name", data)
				input.SetText(string(result.StaticContent))
				// fmt.Println("->->->->->->",result.StaticName + uc.URI().Path())
				filepath = uc.URI().Path();
				w.SetTitle("filePath")
			},w)
		openFileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))

		openFileDialog.Show()
	})



	menuItem := fyne.NewMenu("file",new1,save1,saveAs1,open1)
	menu := fyne.NewMainMenu(menuItem)
	w.SetMainMenu(menu)
	/*
	w.SetContent(
		container.NewWithoutLayout(
			input,
		),
	)
	w.ShowAndRun();
	*/
	notepadContainer := container.NewWithoutLayout(
			input,
		)
	w.SetContent(container.NewBorder(DeskBtn,nil,nil,nil,notepadContainer),
	)
	w.Show()

	/*
	w:=myApp.NewWindow("Calc");
	w.Resize(fyne.NewSize(600,400));
	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calcContainer),
	)
	w.Show();
	*/


}