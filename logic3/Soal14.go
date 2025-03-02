package logic3

import go_print_slice "github.com/Yape0/go-print-slice"

func TigaEmpatbelas(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		num := start
		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 34
		}
		start += 2
	}
	return result
}
