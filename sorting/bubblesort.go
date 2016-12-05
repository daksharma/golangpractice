package sorting

import (
	"fmt"
	"math/rand"
)

// BubbleSort sorts the data to the right side
// heaviest portion trickles down towards the right side
func BubbleSort(original []int, arr []int, nth int) {
	tempArray := arr
	arrLen := len(arr)
	hasSwapped := false
	for count := 0; count < arrLen-1; count++ {
		hasSwapped = false
		for j := 0; j < arrLen-count-1; j++ {
			if tempArray[j] > tempArray[j+1] {
				swapElements(&tempArray[j], &tempArray[j+1])
				hasSwapped = true
			}
		}
		if hasSwapped == false {
			break
		}
	}
	printBubbleSort(original, tempArray)
}

func swapElements(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func printBubbleSort(sorted []int, arr []int) {
	fmt.Printf("***************************\n")
	fmt.Printf("* Unsorted\tSorted\t***\n")
	fmt.Printf("***************************\n")
	for counti := 0; counti < len(arr); counti++ {
		fmt.Println("    ", sorted[counti], "    |\t", arr[counti])
	}
}
