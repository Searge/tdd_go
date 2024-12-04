package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

// TestSum tests the Sum function to ensure it correctly returns the sum of
// all the numbers in the given slice of integers. It checks that the result
// matches the expected output for two different input slices.
func TestSum(t *testing.T) {

	t.Run("collection of 5 numbes", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// TestSumAll tests the SumAll function to ensure it correctly returns a slice
// containing the sums of each provided slice of integers. It checks that the
// result matches the expected output for given input slices.
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sums)
	// Output:
	// [3 9]
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
