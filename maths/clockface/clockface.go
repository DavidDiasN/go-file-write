package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
	minuteHandLength = 80
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinuteHand(w io.Writer, t time.Time) {
	p := makeHand(SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func SecondHand(w io.Writer, t time.Time) {
	p := makeHand(MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func SimpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func TestName(t time.Time) string {
	return t.Format("15:04:05")
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
