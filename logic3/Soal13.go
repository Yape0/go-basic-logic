package logic3

//fix

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaTigabelas(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		for j := 0; j <= mid; j++ {
			if i+j == mid || (i+j)/2 == mid {
				result[j][i] = start
				result[j][n-1-i] = start
				result[n-1-j][i] = start
				result[n-1-j][n-1-i] = start
				start += 2
			}
		}
	}
	return result
}
