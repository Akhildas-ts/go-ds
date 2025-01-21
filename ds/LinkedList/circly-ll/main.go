package main

import (
	"fmt"
	"log"
)

type node struct {
	data int
	next *node
}

type circularList struct {
	head *node
}

func main() {

	l := circularList{}

	l.circularList(10)
	l.circularList(20)
	l.circularList(30)
	l.circularList(40)
	l.circularList(50)

	displayList(l)
}

func (l *circularList) circularList(x int) {

	newNode := &node{data: x}
	if l.head == nil {
		newNode.next = newNode

		l.head = newNode
		return

	}

	current := l.head

	for current.next != l.head {

		current = current.next
	}

	newNode.next = l.head
	current.next = newNode

}

func displayList(l circularList) {

	if l.head == nil {

		log.Println("empty list  ")
		return
	}

	current := l.head

	for {

		fmt.Println(current.data)

		current = current.next

		if current == l.head {
			break
		}
	}

}
