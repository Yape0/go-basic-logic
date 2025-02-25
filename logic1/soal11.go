package logic1

func SatuSebelas(n int) []interface{} {
	slice := make([]interface{}, 0, n)
	start := 3
	slice = append(slice, "buzz")
	slice = append(slice, 1)
	slice = append(slice, "buzz")

	for i := 1; i <= n-3; i++ {
		if i%2 == 0 {
			slice = append(slice, "buzz")
		} else if i%2 == 1 {
			slice = append(slice, start)
			start *= 2
		}
	}
	return slice
}
