package main

import go_print_slice "github.com/Yape0/go-print-slice"

func DuaSatu(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 1 + j*2
		}
	}
	return result
}

func DuaDua(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 2 + j*2
		}
	}
	return result
}

func DuaTiga(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 1 + 2*j + (18 * i)
		}
	}
	return result
}

func DuaEmpat(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		current := start
		count := 0

		for j := 0; j < n; j++ {
			result[i][j] = current

			if count%12 < 6 {
				current += 3
			} else {
				current += 2
			}
			count++
		}
		start += 27
	}
	return result
}

func DuaEnam(n int) [][]int {
	result := go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][n-1-j] = start
				start += 2
			}
		}

	}
	return result

}
