package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("This is our first Array sum test", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func TestSumAll(t *testing.T) {

	t.Run("Test SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test SumAll again", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1})
		want := []int{4, 4, 4, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Test sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func Sum(arr []int) (total int) {
	for _, num := range arr {
		total += num
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, val := range numbersToSum {
		if len(val) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(val[1:])
		}
	}

	return sums
}
