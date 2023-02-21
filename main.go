package main

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const (
	colorOne = "#FF7CCB"
	colorTwo = "#FDFF8C"
	char     = '█'
	width    = 100
)

var now = time.Now()
var dayStart = time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
var dayEnd = time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, time.Local)

// TODO: flags to set color
func main() {
	prog := progress.New(
		progress.WithScaledGradient(colorOne, colorTwo),
		progress.WithColorProfile(termenv.TrueColor),
	)
	prog.Width = width
	prog.ShowPercentage = false
	prog.Full = char
	prog.Empty = char
	percentage := dayProgress()

	start := dayStart.Format(time.Kitchen)
	end := dayEnd.Format(time.Kitchen)

	text := fmt.Sprintf("%.0f%% of day done", math.Round(percentage*100))
	title := lipgloss.NewStyle().Width(width + len(start) + len(end)).Align(lipgloss.Center).Faint(true)
	fmt.Print(title.Render(text))

	fmt.Printf("\n%s %s %s", start, prog.ViewAs(percentage), end)

}

func dayProgress() float64 {

	totalTimeRange := dayEnd.Sub(dayStart).Minutes()
	sinceDayStart := now.Sub(dayStart).Minutes()

	percentage := sinceDayStart / totalTimeRange

	return percentage
}
