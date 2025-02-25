package logic1

func SatuTiga(n int) []int {
	var numbers []int
	for i := 3; i <= n*3; i += 3 {
		numbers = append(numbers, i)
	}
	return numbers

}
