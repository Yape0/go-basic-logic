package logic1

func SatuDua(n int) []int {
	var numbers []int
	for i := 2; i <= n*2; i += 2 {
		numbers = append(numbers, i)
	}
	return numbers

}
