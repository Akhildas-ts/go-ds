package main

import "fmt"

type work struct{}

func main() {

	v := make(map[string]int)

	v["a"] = 10

	m := v["a"]

	fmt.Println(m)

}
