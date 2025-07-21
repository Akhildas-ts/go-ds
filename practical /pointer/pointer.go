package main

import "fmt"

func main() {

	c := 30

	var p *int = &c

	*p = 17

	fmt.Println("update p ", c)

}
func changeValue(num *int) {
	*num = 100

	fmt.Println("num", num)
}

func main() {
	x := 10
	changeValue(&x)
	fmt.Println(x) // Output: 100
}

