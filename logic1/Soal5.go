package logic1

func SatuLima(n int) []int {
	var numbers []int

	for i := 20; i > n/n; i -= 2 {
		numbers = append(numbers, i)
	}
	return numbers

}
