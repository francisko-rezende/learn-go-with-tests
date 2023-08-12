package arraysAndSlices

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers { // range returns i and value
		total += number
	}

	return total
}

func SumAll(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
