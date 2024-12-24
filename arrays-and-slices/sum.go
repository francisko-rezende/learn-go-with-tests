package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range slicesToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
