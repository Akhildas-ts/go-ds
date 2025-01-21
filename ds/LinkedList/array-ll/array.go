package main

import (
	"fmt"
	"log"
)

type node struct {
	Data int
	Next *node
}
type linkedList struct {
	Head   *node
	length int
}

func (l *linkedList) insertion(x int) {

	if l.Head == nil {

		l.Head = &node{Data: x}
		return
	}

	current := l.Head

	for current.Next != nil {

		current = current.Next

	}

	current.Next = &node{Data: x}

}

func display(l *linkedList) {

	if l.Head == nil {

		log.Println("no value ")
		return
	}

	current := l.Head

	for current != nil {

		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {

	l := linkedList{}
	array := []int{2, 5, 8, 2, 9, 2}

	for _, val := range array {

		l.insertion(val)

	}

	display(&l)

}
