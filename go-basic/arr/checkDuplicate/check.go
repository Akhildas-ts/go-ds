package main

// find duplicate element in the array

func findDuplicate(arr []int) bool {

	store := make(map[int]int)

	for _, val := range arr {

		store[val]++
	}

	for _, val := range store {

		if val > 1 {
			return false
		}
	}

	return true

}
