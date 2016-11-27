package main

import "fmt"

func main() {
	fmt.Println("****** Fibonacci Example ******")
	fmt.Println("|----------------------------------|")
	fmt.Println("|   First 40 Fibonacci Sequence|   |")
	fmt.Println("|----------------------------------|")
	defer fmt.Println("****** End Fibonacci Example ******")
	defer fmt.Println("|----------------------------------|")

	// Stay under 50, else it takes longer to calculate higher numbers of Fibonacci
	for fibcount := 0; fibcount < 40; fibcount++ {
		fmt.Printf("| %2d | %12d | %-12d |\n", fibcount, fib(fibcount), fibRecursive(fibcount))
	}
}

func fib(n int) int {
	a, b := 0, 1
	for count := 0; count < n; count++ {
		a, b = b, a+b
	}
	return b
}

func fibRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}
