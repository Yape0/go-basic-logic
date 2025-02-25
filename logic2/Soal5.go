package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaLima(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				result[i][j] = start
				start += 2
			}
		} else {
			for j := 0; j < n; j++ {
				result[i][n-1-j] = start
				start += 2
			}
		}

	}
	return result
}
