package main

import "fmt"

func main() {
	fmt.Println("****** Fibonacci Example ******")
	defer fmt.Println("****** End Fibonacci Example ******")
	for xcount := 0; xcount < 40; xcount++ {
		fmt.Println(fib(xcount), "\t", fibRec(xcount))
	}
}

func fib(n int) int {
	a := 0
	b := 1
	for count := 0; count < n; count++ {
		a = b
		b = a + b
	}
	return b
}

func fibRec(n int) int {
	if n <= 1 {
		return 1
	}
	return fibRec(n-1) + fibRec(n-2)
}
