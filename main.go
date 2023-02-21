package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/muesli/termenv"
)

const (
	colorOne = "#FF7CCB"
	colorTwo = "#FDFF8C"
	char     = '█'
)

var now = time.Now()
var dayStart = time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
var dayEnd = time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, time.Local)

func main() {
	prog := progress.New(
		progress.WithScaledGradient(colorOne, colorTwo),
		progress.WithColorProfile(termenv.TrueColor),
	)
	prog.Width = 100
	prog.ShowPercentage = false
	prog.Full = char
	prog.Empty = char
	percentage := dayProgress()

	start := dayStart.Format(time.Kitchen)
	end := dayEnd.Format(time.Kitchen)
	fmt.Printf("%s %s %s", start, prog.ViewAs(percentage), end)
}

func dayProgress() float64 {

	totalTimeRange := dayEnd.Sub(dayStart).Minutes()
	sinceDayStart := now.Sub(dayStart).Minutes()

	percentage := sinceDayStart / totalTimeRange

	return percentage
}
