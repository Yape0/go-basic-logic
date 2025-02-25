package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaTujuh(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1

	for j := 0; j < n; j++ {
		result[j] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		result[i][i] = start
		start += 2
	}
	return result
}
