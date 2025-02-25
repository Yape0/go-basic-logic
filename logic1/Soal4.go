package logic1

func SatuEmpat(n int) []int {
	var numbers []int
	for i := 19; i >= n/n; i -= 2 {
		numbers = append(numbers, i)
	}
	return numbers

}
