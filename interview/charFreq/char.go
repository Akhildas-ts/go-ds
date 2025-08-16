package main

import "fmt"

func main() {
	// find char frequency in a string

	s := "hello world"

	r := []rune(s)
	m := make(map[string]int)

	for _, val := range r {
		if val == ' ' {
			continue
		}
		m[string(val)]++
	}
	fmt.Println(m)

}
