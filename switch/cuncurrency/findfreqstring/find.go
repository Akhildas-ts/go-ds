// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// with go routine cuncurreny find the word frequency
func main() {

	wrd := "hello its akhildas ts"
	ch := make(chan map[string]int)

	go findfreq(ch, wrd)

	fmt.Println(<-ch)

}

func findfreq(ch chan map[string]int, wrd string) {
	result := make(map[string]int)
	for i := 0; i < len(wrd); i++ {
		if string(wrd[i]) != " " {
			result[string(wrd[i])]++
		}
	}
	ch <- result
	close(ch)

}
