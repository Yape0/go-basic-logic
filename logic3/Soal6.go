package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaEnam(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j <= n-1 && i <= j {
				if i%2 == 0 {
					result[i][j] = start
					result[n-1-i][j] = start
				} else {
					result[i][n-1-j] = start
					result[n-1-i][n-1-j] = start
				}
				start += 2
			}
		}
	}

	return result
}
