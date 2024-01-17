package clockface_test

import (
	"main/maths/clockface"
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clockface.SimpleTime(0, 0, 30), clockface.Point{0, -1}},
		{clockface.SimpleTime(0, 0, 45), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(clockface.TestName(c.time), func(t *testing.T) {
			got := clockface.SecondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {

}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshhold = 1e-7
	return math.Abs(a-b) < equalityThreshhold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
