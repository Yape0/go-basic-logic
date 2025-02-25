package logic1

func SatuSembilan(n int) []int {
	var numbers []int
	start := 3
	median := 0.5 * float64(n)

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, start)
		if float64(i) < median {
			start += 3
		} else if float64(i) > median {
			start -= 3
		}
	}
	return numbers
}
