package main

import "fmt"

// find duplicate in the array
func duplicate(arr [5]int) [5]int {
	store := make(map[int]int)
	var duplicate [5]int

	for _, val := range arr {

		store[val]++
	}
	i := 0
	for k, v := range store {
		if v > 1 {

			duplicate[i] = k
			i++
		}
	}

	return duplicate

}
func main() {

	dup := duplicate([5]int{5, 6, 5, 2, 2})

	fmt.Println(dup)

}
