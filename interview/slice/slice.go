package main

import "fmt"

// what is the values if a,b and c if it print in the last

func main() {

	//cap interview questions
	a := []int{10, 2, 3, 4, 5, 6}
	//cap(a) = 6

	b := a[1:4]
	//cap(b) = 5

	c := a[:cap(b)]
	//cap(c)= 5
	c[0] = 22
	b[0] = 23

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
