package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaLima(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	mid := (n - 1) / 2
	start := 1

	for i := 0; i <= mid; i++ {

		for j := 0; j <= i; j++ {
			if i >= j && i+j <= n-1 {
				if i%2 == 0 {
					result[i][j] = start
					result[i][n-1-j] = start
					result[i][i-j] = start

				} else {
					result[i][i-j] = start
					result[i][n-1-(i-j)] = start
				}
				start += 2
			}
			for i := mid + 1; i < n; i++ {
				for j := 0; j < n; j++ {
					result[i][j] = result[n-1-i][j]
				}
			}
		}

	}
	return result
}
