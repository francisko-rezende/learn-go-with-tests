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

func SumAllTails(numbersToSum ...[]int) []int {
	var tailsSums []int

	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			tailsSums = append(tailsSums, 0)
		} else {
			tail := numbers[1:]
			tailsSums = append(tailsSums, Sum(tail))
		}

	}
	return tailsSums
}
