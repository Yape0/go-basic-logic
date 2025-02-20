package main

func Satusatu(n int) []int {
	var numbers []int
	for i := 1; i < n*2; i += 2 {
		numbers = append(numbers, i)
	}
	return numbers
}

func SatuDua(n int) []int {
	var numbers []int
	for i := 2; i <= n*2; i += 2 {
		numbers = append(numbers, i)
	}
	return numbers

}

func SatuTiga(n int) []int {
	var numbers []int
	for i := 3; i <= n*3; i += 3 {
		numbers = append(numbers, i)
	}
	return numbers

}

func SatuEmpat(n int) []int {
	var numbers []int
	for i := 19; i >= n/n; i -= 2 {
		numbers = append(numbers, i)
	}
	return numbers

}

func SatuLima(n int) []int {
	var numbers []int

	for i := 20; i > n/n; i -= 2 {
		numbers = append(numbers, i)
	}
	return numbers

}

func SatuEnam(n int) []int {
	var numbers []int

	for i := n * 3; i >= n*3/10; i -= 3 {
		numbers = append(numbers, i)
	}
	return numbers
}

func SatuTujuh(n int) []int {
	var numbers []int
	start := 1
	median := 0.5 * float64(n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, start)
		if float64(i) < median {
			start += 2
		} else if float64(i) > median {
			start -= 2
		}

	}
	return numbers

}

func SatuDelapan(n int) []int {
	var numbers []int
	start := 2
	median := 0.5 * float64(n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, start)
		if float64(i) < median {
			start += 2
		} else if float64(i) > median {
			start -= 2
		}
	}
	return numbers
}

func SatuSembilan(n int) []int {
	var numbers []int
	start := 3
	median := 0.5 * float64(n)

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, start)
		if float64(i) < median {
			start += 3
		} else if float64(i) > median {
			start -= 3
		}
	}
	return numbers
}

func SatuDuabelas(n int) []int {
	var numbers []int
	for i := 0; i < n-9; i++ {
		start := 1
		for j := 0; j < n-8; j++ {
			numbers = append(numbers, start)
			start += 2
		}
		numbers = append(numbers, start)
	}
	return numbers
}
