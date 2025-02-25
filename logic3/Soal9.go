package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaSembilan(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	mid := (n - 1) / 2

	for k := 0; k <= mid; k++ {
		start := 1
		for b := 0; b < mid+1; b++ {
			if k+b >= mid && b-k <= mid {
				result[b][k] = start
				result[b][n-1-k] = start
				result[n-1-b][k] = start
				result[n-1-b][n-1-k] = start
				start += 2
			}

		}

	}
	return result
}
