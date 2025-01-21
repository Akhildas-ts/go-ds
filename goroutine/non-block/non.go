package main

import (
	"fmt"
)

func main() {
	a := 10
	hi(a)

}

func hi(a int) {
	for i := 1; i < a; i++ {

		fmt.Println(i)
		if i == 7 {
			return
		}
	}

	
}
