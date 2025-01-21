package main

import (
	"fmt"
	"log"
)

type node struct {
	Data int
	Next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) insertion(x int) {

	if l.head == nil {
		l.head = &node{Data: x}

		return

	}

	current := l.head

	for current.Next != nil {

		current = current.Next

	}

	current.Next = &node{Data: x}
	l.length++
}

func display(l linkedlist) {

	if l.head == nil {

		log.Println("given list is nil - cound'nt print ")
		return
	}

	current := l.head

	for current != nil {

		fmt.Println(current.Data)
		current = current.Next
	}

}

func (l *linkedlist) insertAfterelement(x, i int) {

	if l.head == nil {

		log.Println("linked list empty..")
		return

	}

	newNode := &node{Data: x}

	current := l.head

	for current.Next != nil {

		if current.Data == i {

			newNode.Next = current.Next
			current.Next = newNode
			return
		}

		current = current.Next

	}

}

func (l *linkedlist) delete(x int) {

	if l.head == nil {

		log.Println("linked is empty - error ")
		return

	}
	current := l.head
	prev := current

	for current.Next != nil {

		if current.Data == x {

			prev.Next = current.Next
			current = prev
			return

		}
		prev = current
		current = current.Next

	}
}

func main() {

	l := linkedlist{}

	l.insertion(10)
	l.insertion(20)
	l.insertion(30)
	l.insertion(40)
	l.insertion(50)
	l.insertion(60)

	l.insertAfterelement(35, 30)

	display(l)

}
