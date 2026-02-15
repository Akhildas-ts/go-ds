package main

import "fmt"

func main() {
	m := make(map[string]struct {
		age    int
		gender string
	})
	m["akhil"] = struct {
		age    int
		gender string
	}{
		age:    20,
		gender: "male",
	}
	fmt.Println(m)
}
