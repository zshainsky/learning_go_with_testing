package main

// Sum function takes in a slice of ints, sums all of the integers in that slice and returns sum value.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes in a variable number of slices of integers and returns the sum of each slice in a slice of integers.
func SumAll(numbersToSum ...[]int) []int {
	// Create a slice
	var sums []int

	// Append sum values to slice
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// Create a slice
	var sums []int

	// Loop through all slices in numbersToSum
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Sum all values in `numbers` skipping the first index ("head")
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums

}
