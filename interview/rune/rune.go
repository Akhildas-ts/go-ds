package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "hello, World!"

	r := []rune(s)

	for i := 0; i < len(r); i++ {

		if i == 0 {

			upper := unicode.ToUpper(r[i])
			r[i] = upper
		} else if string(r[i]) == " " {
			upper := unicode.ToUpper(r[i+1])

			r[i+1] = upper
		}

	}
	fmt.Println(string(r))

}
