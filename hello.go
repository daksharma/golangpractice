package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(fib(2))
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
