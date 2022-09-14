package algo

// Fibonacci returns the n-th fibonacci number.
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	// fibo(n) = fibo(n-1) + fibo(n-2)
	// n1 : fibo(n-1)
	// n2 : fibo(n-2)
	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
