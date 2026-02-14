package main

//finding anagrams in a list of strings

import "fmt"

// 
func main() {

	arr := []string{"eat", "bat", "ate", "tea", "tab"}
	result := [][]string{}

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {

			f := arr[i]
			s := arr[j]
			count := 0
			if len(f) != len(s) {

				continue


			}
			mini := []string{}
			for l := 0; l < len(f); l++ {

				for k := 0; k < len(s); k++ {

					if f[l] == s[k] {

						count++
						break
					}
				}

			}

			if count == len(f) {

				mini = append(mini, f, s)
				result = append(result, mini)
			}

		}

	}

	fmt.Println(result)

}