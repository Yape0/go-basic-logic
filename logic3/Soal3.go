package logic3

//fix

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaTiga(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				if i == 0 {
					result[i][j-i] = start
					start += 3
				} else if i%2 == 0 {
					result[i][j-i] = start
					start += 2
					if j == n-1 {
						start += 1
					}
				} else {
					result[i][n-j-1] = start
					start += 3
					if j == n-1 {
						start += 1
					}
				}
			}

		}
	}

	return result
}
