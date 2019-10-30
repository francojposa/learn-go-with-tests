package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		want := 6
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		if got != want {
			t.Errorf("\nwant: %d, got: %d, given: %v", want, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	want := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{0, 9})

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("Sum the tails of non-empty slices", func(t *testing.T) {
		want := []int{10, 4}
		got := SumAllTails([]int{3, 4, 6}, []int{1, -3, 7})

		checkSums(t, got, want)
	})

	t.Run("Sum tails of empty & non-empty slices", func(T *testing.T) {
		want := []int{0, 9}
		got := SumAllTails([]int{}, []int{3, 4, 5})

		checkSums(t, got, want)
	})

}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwant %v, got %v}", want, got)
	}
}
