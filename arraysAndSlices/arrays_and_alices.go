package arraysAndSlices

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers { // range returns i and value
		total += number
	}

	return total
}
