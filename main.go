package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03 H 04 M 05 Seconds")
	clock.SetText(formatted)
}

func getSecondOfDay(t time.Time) int {
    return 60*60*t.Hour() + 60*t.Minute() + t.Second()
}

func main() {
	a := app.New()
	w := a.NewWindow("Display Time")

	clock := widget.NewLabel(" ")
	clock.Alignment = fyne.TextAlignCenter
	clock.TextStyle	= fyne.TextStyle{Bold: true, Monospace: true}
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second){
			updateTime(clock)
		}
	}()
	
	w.SetContent(clock)
	
	w.Resize(fyne.NewSize(300, 100))
	w.Show()	
	a.Run()
}