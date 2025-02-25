package logic3

//fix

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaEmpatGagal(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 0 {
					result[i][j] = start
				} else {
					result[i][i-j] = start
				}
				start += 2
			}
		}
	}
	return result
}

func TigaEmpat(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 0 {
					result[i][n-1-j] = start
				} else {
					result[i][j] = start

				}
				start += 2
			}
		}
	}
	return result
}

//result[b][k] = start
//result[b][n-1-k] = start
//result[n-1-b][k] = start
//result[n-1-b][n-1-k] = start
