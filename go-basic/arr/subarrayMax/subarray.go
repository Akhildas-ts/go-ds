package main

import (
	"fmt"
)

//find max sum value in subarray

func main() {

	arr := []int{2, -1, 8, -9, 5, 2, 4}

	currentSum := 0
	max := -100000

	for _, val := range arr {

		currentSum += val

		if max < currentSum {

			max = currentSum
		}

		if currentSum < 0 {

			currentSum = 0
		}

	}
	// here we get the max value
	fmt.Println(max)
}
