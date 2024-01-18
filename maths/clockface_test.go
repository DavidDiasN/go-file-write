package clockface_test

import (
	"main/maths/clockface"
	"math"
	"testing"
	"time"
)

// func TestSecondHandAtMidnight(t *testing.T) {
// 	cases := []struct {
// 		time  time.Time
// 		point clockface.Point
// 	}{
// 		{clockface.SimpleTime(0, 0, 30), clockface.Point{0, -1}},
// 		{clockface.SimpleTime(0, 0, 45), clockface.Point{-1, 0}},
// 	}

// 	for _, c := range cases {
// 		t.Run(clockface.TestName(c.time), func(t *testing.T) {
// 			got := clockface.SecondHandPoint(Stdout, c.time)
// 			if !roughlyEqualPoint(got, c.point) {
// 				t.Fatalf("wanted %v Point, but got %v", c.point, got)
// 			}
// 		})
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {

// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 + 90}
// 	got := clockface.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshhold = 1e-7
	return math.Abs(a-b) < equalityThreshhold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

// func TestSVGWriterAtMidnight(t *testing.T) {
// 	tm := time.Date(1336, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	var b strings.Builder
// 	clockface.SVGWriter(&b, tm)
// 	got := b.String()

// 	want := `<line x0="150" y1="150" x2="150" y2="60"`

// 	if !strings.Contains(got, want) {
// 		t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
// 	}
// }

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{clockface.SimpleTime(0, 30, 0), math.Pi},
		{clockface.SimpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(clockface.TestName(c.time), func(t *testing.T) {
			got := clockface.MinutesInRadians(c.time)
			if got != c.angle {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clockface.SimpleTime(0, 30, 0), clockface.Point{0, -1}},
		{clockface.SimpleTime(0, 45, 0), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(clockface.TestName(c.time), func(t *testing.T) {
			got := clockface.MinuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

