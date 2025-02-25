package logic1

func SatuDelapan(n int) []int {
	var numbers []int
	start := 2
	median := 0.5 * float64(n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, start)
		if float64(i) < median {
			start += 2
		} else if float64(i) > median {
			start -= 2
		}
	}
	return numbers
}
