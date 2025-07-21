package main

import "fmt"

func TripleArr(arr []int, target int) [][]int {

	result := [][]int{}

	// Iterate through all possible triplets
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == target {
					// Add the triplet to the result
					result = append(result, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}

	return result

}
func main() {

	arr := []int{0, -1, 2, -3, 1}
	fmt.Println(TripleArr(arr, -2))

}
