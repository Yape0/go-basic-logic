package logic1

func Satusatu(n int) []int {
	var numbers []int
	for i := 1; i < n*2; i += 2 {
		numbers = append(numbers, i)
	}
	return numbers
}
