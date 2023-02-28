package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const (
	colorOne     = "#FF7CCB"
	colorTwo     = "#FDFF8C"
	char         = '█'
	defaultWidth = 100
)

var now = time.Now()
var dayStart = "9:30AM"
var dayEnd = "5:30PM"

func main() {
	colorLeft := flag.String("color-1", colorOne, "Left color")
	colorRight := flag.String("color-2", colorTwo, "Right color")
	width := flag.Int("width", defaultWidth, "width")
	start := flag.String("start", dayStart, "start time")
	end := flag.String("end", dayEnd, "end time")
	flag.Parse()

	dayStart := parseTimeFlag(*start)
	dayEnd := parseTimeFlag(*end)

	prog := progress.New(
		progress.WithScaledGradient(*colorLeft, *colorRight),
		progress.WithColorProfile(termenv.TrueColor),
	)
	prog.Width = *width
	prog.ShowPercentage = false
	prog.Full = char
	prog.Empty = char
	percentage := dayProgress(dayStart, dayEnd)

	startKitchen := dayStart.Format(time.Kitchen)
	endKitchen := dayEnd.Format(time.Kitchen)

	s := math.Round(percentage * 100)
	if s > 100 {
		s = 100
	}

	text := fmt.Sprintf("%.0f%% of day done", s)
	title := lipgloss.NewStyle().Width(defaultWidth + len(startKitchen) + len(endKitchen)).Align(lipgloss.Center).Faint(true)
	fmt.Print(title.Render(text))

	fmt.Printf("\n%s %s %s", startKitchen, prog.ViewAs(percentage), endKitchen)

}

func parseTimeFlag(str string) time.Time {
	layout := time.Kitchen
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Printf("Could not parse time %s\nError: %v\n", str, err)
	}

	return t
}

func dayProgress(s time.Time, e time.Time) float64 {
	start := time.Date(now.Year(), now.Month(), now.Day(), s.Hour(), s.Minute(), 0, 0, time.Local)
	end := time.Date(now.Year(), now.Month(), now.Day(), e.Hour(), e.Minute(), 0, 0, time.Local)

	totalTimeRange := end.Sub(start).Minutes()
	sinceDayStart := now.Sub(start).Minutes()

	percentage := sinceDayStart / totalTimeRange

	return percentage
}
