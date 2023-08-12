package arraysAndSlices

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers { // range returns i and value
		total += number
	}

	return total
}

func SumAll(numberToSum ...[]int) []int {
	lengthOfNumbers := len(numberToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
