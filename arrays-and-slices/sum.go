package arraysandslices

func Sum(array [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += array[i]
	}
	return sum
}
