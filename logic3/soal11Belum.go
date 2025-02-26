package logic3

//belom selesai

//func TigaSebelas(n int) [][]int {
//	result := go_print_slice.CreateSlice(n)
//	mid := (n - 1) / 2
//	start := 1
//	//num := 1
//
//	for i := 0; i <= mid; i++ {
//		for j := n - 1; j >= mid; j-- {
//			if i == j {
//				result[i][j] = start
//				result[n-1-i][j] = start
//			}
//			start += 2
//		}
//		start = 1
//		for i := 0; i <= mid; i++ {
//			if i%2 == 0 {
//				for j := n - 1; j >= mid; j-- {
//					if i+j >= n-1 {
//						result[i][j] = start
//						result[n-1-i][j] = start
//						start += 2
//					} else {
//						for j := n - 1; j >= mid; j-- {
//							if i+j >= n-1 {
//								result[i][(2*(n-1) - i - j)] = start
//								result[i][2*(n-1)-i-j] = start
//							}
//						}
//					}
//				}
//
//			}
//
//		}
//	}
//	return result
//}

//func TigaSebelas(n int) [][]int {
//	result := go_print_slice.CreateSlice(n)
//	start := 1
//	mid := (n - 1) / 2
//
//	for i := 0; i <= mid; i++ {
//		for j := 0; j <= mid; j++ {
//			if i == j {
//				result[i][j] = start
//				result[n-1-i][j] = start
//			}
//		}
//		start += 2
//	}
//	start = 1
//	for i := 0; i <= mid; i++ {
//		for j := 0; j <= mid; j++ {
//			if i >= j {
//				if i
//				result[i][n-1] = start
//				result[n-1-i][n-1-j] = start
//			}
//		}
//	}
//	return result
//}
