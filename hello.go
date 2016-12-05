package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/daksharma/golangtest/recursion"
	"github.com/daksharma/golangtest/sorting"
)

const numSize = 50

func main() {
	fmt.Println("Hello World!")
	recursion.FibonacciNums()
	fmt.Println("Bubble Sort")
	randomNums := make([]int, numSize)
	originals := make([]int, numSize)

	for count := 0; count < numSize; count++ {
		// fmt.Println(count)
		rand.Seed(int64(time.Now().Nanosecond()))
		randVal := rand.Int()
		randomNums[count] = obtainSmallNumber(randVal, 2)
		originals[count] = obtainSmallNumber(randVal, 2)
	}

	sorting.BubbleSort(originals, randomNums, len(randomNums))
}

func obtainSmallNumber(n int, size int) int {
	intStr := strconv.Itoa(n)  // string version
	smallInt := intStr[0:size] // certain range from the string version
	nRes, _ := strconv.Atoi(smallInt)
	return nRes // return the range version as an int
}
