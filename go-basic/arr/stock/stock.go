package main

import "fmt"

// here im writing the code for a stock

// we buy a product from at good time
//day 0: 7
//day1: 4
//day2: 5
//day3: 1
//day4: 2
//day5: 5
// here we can buy the product at day3 at 1 this the best offer
// then we can sale at day 5 so this was the best price
// so in this case our profit is 4, 1 - 5 == 4

func main() {

	arr := []int{7, 4, 5, 1, 2, 5}

	fmt.Println(profit(arr))

}

func profit(arr []int) int {

	// getting the product on low time

	low := arr[0]
	index := 0
	high := 0

	for i := 0; i < len(arr); i++ {

		if low > arr[i] {
			low = arr[i]
			index = i
		}
	}

	for j := index + 1; j < len(arr); j++ {

		if high < arr[j] {

			high = arr[j]

		}
	}

	result := 0

	if high == 0 {
		return result
	}

	if low < high {

		result = high - low

	}

	return result

}
