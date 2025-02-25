package logic1

func SatuEnam(n int) []int {
	var numbers []int

	for i := n * 3; i >= n*3/10; i -= 3 {
		numbers = append(numbers, i)
	}
	return numbers
}
