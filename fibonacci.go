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
	for xcount := 0; xcount < 40; xcount++ {
		fmt.Printf("| %2d | %12d | %-12d |\n", xcount, fib(xcount), fibRec(xcount))
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
