package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(0, 0, 30), math.Pi},
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{SimpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {

		t.Run(TestName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)

			if got != c.angle {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}

		})
	}
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}

	SVGWriter(&b, tm)

	svg := SVG{}

	xml.Unmarshal(b.Bytes(), &svg)

	want := Line{150, 150, 150, 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			SimpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			SimpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			SimpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}

		})
	}
}

// func TestSVGWriterHourHand(t *testing.T) {
// 	cases := []struct {
// 		time time.Time
// 		line Line
// 	}{
// 		{
// 			SimpleTime(6, 0, 0),
// 			Line{150, 150, 150, 200},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(TestName(c.time), func(t *testing.T) {
// 			b := bytes.Buffer{}
// 			SVGWriter(&b, c.time)

// 			svg := SVG{}
// 			xml.Unmarshal(b.Bytes(), &svg)

// 			if !containsLine(c.line, svg.Line) {
// 				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
// 			}
// 		})
// 	}
// }

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(6, 0, 0), math.Pi},
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(21, 0, 0), math.Pi * 1.5},
		{SimpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := HoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshhold = 1e-7
	return math.Abs(a-b) < equalityThreshhold
}

// func roughlyEqualPoint(a, b Point) bool {
// 	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
// }
