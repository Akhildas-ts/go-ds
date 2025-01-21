package main

import "fmt"

func longestSubString(s string) int {

	longest := 0
	word := ""

	for i := 0; i < len(s); i++ {

		if string(s[i]) == " " {

			if longest < len(word) {

				longest = len(word) - 1

			}

			word = ""
		}

		word += string(s[i])

	}

	return longest

}
func main() {

	fmt.Println(longestSubString("hello my name is akhil from"))
}
