package logic1

func SatuSepuluh(n int) []interface{} {
	slice := make([]interface{}, 0, n)
	start := 2
	fizz := "fizz"
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			slice = append(slice, start)
			start *= 2
		} else if i%2 == 1 {
			slice = append(slice, fizz)
		}

	}
	return slice
}
