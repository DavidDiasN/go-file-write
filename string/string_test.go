package string

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("This is a string compare test expecting true", func(t *testing.T) {
		got := strings.Compare("david", "david")
		want := 0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing strings.Contains", func(t *testing.T) {
		got := strings.Contains("the crowd goes wild", "wild")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Test ContainsAny", func(t *testing.T) {
		got := strings.ContainsAny("david", "d")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Test cut", func(t *testing.T) {
		got1, got2, gotbool := strings.Cut("Hello world!", "wo")
		want1, want2, wantbool := "Hello ", "rld!", true
		if got1 != want1 || got2 != want2 || gotbool != wantbool {
			t.Errorf("got1 %s want1 %s got2 %s want2 %s gotbool %t wantbool %t", got1, want1, got2, want2, gotbool, wantbool)
		}
	})

	t.Run("CutPrefix", func(t *testing.T) {
		got, gotbool := strings.CutPrefix("log: 2018, 3000, 5000", "log: ")
		want, wantbool := "2018, 3000, 5000", true
		if want != got || gotbool != wantbool {
			t.Errorf("got %s want %s gotbool %t wantbool %t", got, want, gotbool, wantbool)
		}
	})

	t.Run("Testing Cut Suffic", func(t *testing.T) {
		got, gotBool := strings.CutSuffix("message coming through over", " over")
		want, wantBool := "message coming through", true
		if got != want || gotBool != wantBool {
			t.Errorf("got %s want %s gotBool %t wantBool %t", got, want, gotBool, wantBool)
		}
	})

	t.Run("Testing EqualFold case sensitify ingoring function", func(t *testing.T) {
		// this should return true
		got := strings.EqualFold("david", "DaviD")
		if !got {
			t.Error("got was not true, but it should be")
		}
	})

	t.Run("tseing fields functions of the strings class wooo", func(t *testing.T) {
		got := strings.Fields("It's the best day ever")
		want := []string{"It's", "the", "best", "day", "ever"}
		for index, val := range got {
			if val != want[index] {
				t.Errorf("got %v want %v", got, want[index])
			}
		}

	})

	t.Run("checking the Index method", func(t *testing.T) {
		got := strings.Index("David is wow so cool", "wow")
		want := 9
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
