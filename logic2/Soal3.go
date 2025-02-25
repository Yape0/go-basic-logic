package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaTiga(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 1 + 2*j + (18 * i)
		}
	}
	return result
}
