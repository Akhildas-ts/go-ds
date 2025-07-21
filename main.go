package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 14

	str := strconv.Itoa(num)

	for _, val := range str {
		fmt.Println(string(val))
	}

}
