// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func selection(arr []int) {

	for i := 0; i < len(arr)-1; i++ {

		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j

			}
		}
		arr[i], arr[min] = arr[min], arr[i]

	}
}

func main() {

	arr := []int{5, 3, 4, 2}
	selection(arr)
	fmt.Println(arr)
	fmt.Println("Hello, 世界")
}
