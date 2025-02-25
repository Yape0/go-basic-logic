package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaDua(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		num := 2

		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}
	}
	return result
}
