package main

import "fmt"

// given a array rearange with max min second max in frist

func main() {

	input := []int{10, 2, 3, 4, 5, 6, 7, 8}
	//output = 9,1,7,2,3,4,5,6

	fmt.Println(rearrange(input))

}

func rearrange(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	max := 0
	min := 10000
	secondMax := 0

	for _, val := range arr {

		if max < val {
			secondMax = max
			max = val
		} else if secondMax < val {
			secondMax = val
		}
		if min > val {
			min = val
		}

	}

	fmt.Println(max, min, secondMax)
	seen[max] = true
	seen[secondMax] = true
	seen[min] = true

	result = append(result, max, min, secondMax)

	for _, val := range arr {

		if !seen[val] {
			result = append(result, val)
		}
	}

	return result
}
