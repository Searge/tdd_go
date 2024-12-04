package arrays

// Sum takes a slice of integers and returns the sum of all the numbers in the slice.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes in a variadic number of slices of integers and returns a slice of
// the sums of each slice.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(head int, numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[head:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
