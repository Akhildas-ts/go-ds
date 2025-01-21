package main

import "fmt"

type node struct {
	Data int
	Next *node
	Prev *node
}

type doublylist struct {
	Head *node
	Tail *node
}

func (d *doublylist) doublylistInsertion(x int) {

	addNode := &node{Data: x}

	if d.Head == nil {

		d.Head = addNode
		d.Tail = addNode
		return

	}

	current := d.Head

	for current.Next != nil {

		current = current.Next
	}

	addNode.Prev = current
	current.Next = addNode
	d.Tail = addNode

}

func doublylistPrint(d *doublylist) {

	current := d.Head

	for current.Next != nil {

		fmt.Println(current.Data)

		current = current.Next
	}

}

func main() {

	d := doublylist{}
	d.doublylistInsertion(10)
	d.doublylistInsertion(20)

	d.doublylistInsertion(30)
	d.doublylistInsertion(40)
	d.doublylistInsertion(50)
	d.doublylistInsertion(60)

	doublylistPrint(&d)
	

}
