package main

import "fmt"

// giving a slice add a element without overwriting the existing element
func main() {

	s := []int{1, 2, 4, 5, 6}
	index := 2

	s = append(s[:index], append([]int{3}, s[index:]...)...)

	fmt.Println(s)

}
