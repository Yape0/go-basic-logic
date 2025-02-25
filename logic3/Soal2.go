package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaDua(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][n-1-(j-i)] = start
				start += 2
			}
		}
	}
	return result
}
