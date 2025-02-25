package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaDuabelas(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 1
		for j := 0; j < n; j++ {
			if j >= 1 && j <= n-1 {
				result[i][j] = start
				start += 2
			} else if j >= n-i-1 && j <= i {
			}
		}
	}
	return result
}
