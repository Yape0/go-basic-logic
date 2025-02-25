package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaEmpat(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		current := start
		count := 0

		for j := 0; j < n; j++ {
			result[i][j] = current

			if count%12 < 6 {
				current += 3
			} else {
				current += 2
			}
			count++
		}
		start += 27
	}
	return result
}
