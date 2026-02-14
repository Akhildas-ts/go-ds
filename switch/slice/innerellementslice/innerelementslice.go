package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	//here new array point to the same array, starting from index 1 to index 4
	//len will be 4-1 = 3
	//cap will be original cap(5) - low index(1) = 4
	s1c1 := arr[1:4]

	//second slice from the first slice
	//here new slice point to the s1c1, s1c1 pointing to the arr
	//so when we print slc2 it will print the element from arr starting from index 1 to index 4
	slc2 := s1c1[:4]

	fmt.Println("len", len(s1c1), "cap", cap(s1c1), "slice", s1c1)

	fmt.Println("len", len(slc2), "cap", cap(slc2), "slice", slc2)

	//inner element slice

	// slc := []int{1, 2, 3, 4, 5}

	// slc = slc[1:3]
	// //here len will low(1) - high(3) = 2
	// //cap will low(1)- origial cap(5) = 4
	// // why in the case of cap we are not doing high, because cap is calculated from the low index to the end of the original slice

	// fmt.Println("len", len(slc), "cap", cap(slc), "slice", slc)

}
