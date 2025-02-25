package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaDelapan(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	mid := (n - 1) / 2
	start := 1

	for k := 0; k <= mid; k++ {
		for b := 0; b < n; b++ {
			if k+b >= mid && b-k <= mid {
				if k%2 == 1 {
					result[b][k] = start
					result[b][n-1-k] = start
				} else {
					result[n-1-b][k] = start
					result[n-1-b][n-1-k] = start
				}
				start += 2
			}

		}

	}
	return result
}
