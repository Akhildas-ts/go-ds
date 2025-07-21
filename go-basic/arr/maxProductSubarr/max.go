package main

import "fmt"

func main() {

	arr := []int{1, 4, -1, 6, 6, -5, 8, 9}

	fmt.Println(maxProductSub(arr))

}

func maxProductSub(arr []int) int {

	max := -10000
	currentSum := 1

	for i, val := range arr {

		currentSum *= val

		fmt.Println("i", i, "cuurentSum", currentSum)

		if currentSum > max {

			max = currentSum
		}

		if currentSum < 0 {

			currentSum = 1
		}

	}

	return max
}
