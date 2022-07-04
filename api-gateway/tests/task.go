package tests

func Fib3(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib3(n-1) + Fib3(n-2)
	}
}

