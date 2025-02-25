package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaSebelas(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 17
		for j := 8; j >= i; j-- {
			result[i][j] = start
			start -= 2
		}
	}
	return result
}
