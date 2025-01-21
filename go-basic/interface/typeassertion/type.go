package main

import "fmt"

type person struct{}

func (p person) greet() {

	fmt.Println("hello its akhil")
}

func main() {

	p := person{}
	p.greet()

}
