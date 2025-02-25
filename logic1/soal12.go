package logic1

func SatuDuaBelas(n int) []int {
	slice := make([]int, n)
	start := 1

	for i := 0; i < n; i++ {
		slice[i] = start
		start += 2
		if start > 7 {
			start = 1
		}
	}
	return slice
}
