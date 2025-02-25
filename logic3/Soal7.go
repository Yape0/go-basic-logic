package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaTujuh(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	mid := (n - 1) / 2
	start := 1

	for row := 0; row <= mid; row++ {
		for col := 0; col < n; col++ {
			if row+col >= mid && col-row <= mid {
				if row%2 == 1 {
					result[row][col] = start
					result[n-1-row][col] = start
				} else {
					result[row][n-1-col] = start
					result[n-1-row][n-1-col] = start
				}
				start += 2
			}
		}
	}
	return result
}
