package main

import "fmt"

//second largest array

func main() {

	largest := 0
	secondlargest := 0

	arr := []int{12, 4, 8, 5, 27, 9}

	for _, val := range arr {

		if largest < val {
			secondlargest = largest
			largest = val
		}
	}

	fmt.Println(secondlargest)
	fmt.Println(largest)

}
