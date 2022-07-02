package main

import "fmt"

func fib(n int) int {
	fibonacci := make([]int, n+1)

	if n < 2 {
		return n
	}

	fibonacci[0] = 0
	fibonacci[1] = 1

	for i := 2; i <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci[n]
}

func main() {
	fmt.Printf("%v\n", fib(3))
}
