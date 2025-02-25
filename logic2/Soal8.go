package logic2

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaDelapan(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 17
	count := 8

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		result[j][count] = start
		start -= 2
		count--

	}
	return result
}
