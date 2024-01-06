package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Testing Repeat functionality", func(t *testing.T) {
		repeated := Repeat(5, "a")
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Test for adding int", func(t *testing.T) {
		repeated := Repeat(7, "a")
		expected := "aaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func Repeat(rep int, let string) (repeatString string) {
	for i := 0; i < rep; i++ {
		repeatString += let
	}
	return repeatString
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(6, "a")
	}
}

func ExampleRepeat() {

	sum := Repeat(6, "d")
	fmt.Println(sum)
	// Output: dddddd
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
