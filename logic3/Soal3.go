package logic3

//fix

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaTiga(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j <= n-1 {
				if i%2 == 0 {
					result[i][j] = start
					start += 2
				} else {
					result[i][i-j] = start
					start += 3
				}
			}

		}
	}

	return result
}
