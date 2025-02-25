package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaSembilan(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	count := 8

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		result[i][i] = start
		result[count][i] = start
		start += 2
		count--
	}
	return result
}
