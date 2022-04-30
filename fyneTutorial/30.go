package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)
func main(){
	a := app.New();
	w := a.NewWindow("multiline entry");
	w.Resize(fyne.NewSize(600, 400));

	multilineEntry := widget.NewMultiLineEntry();
	const (
		loremIpsum = "Lorem Ipsum refers to a dummy block of text that is often used in publishing and graphic design to fill gaps in the page before the actual words are put into the finished product. Lorem ipsum resembles Latin but has no real meaning."
	);
	multilineEntry.SetText(loremIpsum);
	// multilineEntry.Wrapping = fyne.TextWrapBreak;
	// multilineEntry.Wrapping = fyne.TextWrapOff;
	multilineEntry.Wrapping = fyne.TextTruncate;
	w.SetContent(multilineEntry);


	w.ShowAndRun();
}