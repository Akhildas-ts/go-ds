package main

import "fmt"

type Listnode struct {
	data int
	next *Listnode
}

func mergeTwoLists(list1 *Listnode, list2 *Listnode) *Listnode {

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	dummy := &Listnode{}

	current := dummy
	for list1 != nil && list2 != nil {

		if list1.data > list2.data {

			current.next = list2
			list2 = list2.next

		} else if list1.data < list2.data {

			current.next = list1
			list1 = list1.next

		}

		current = current.next

	}

	if list1 == nil {

		current.next = list2
	} else {

		current.next = list1
	}

	return dummy.next
}
func (l *Listnode) insertNode(data int) {

	if l == nil {

		l.data = data
		return
	}

	current := l

	for current.next != nil {

		current = current.next
	}

	current.next = &Listnode{data: data}

}

func display(l *Listnode) {

	if l == nil {
		fmt.Println("node is nil ")

		return
	}

	for l != nil {

		fmt.Println(l.data)

		l = l.next
	}
}

func main() {

	l := Listnode{}

	l2 := Listnode{}

	l2.insertNode(1)
	l2.insertNode(2)
	l2.insertNode(3)
	l2.insertNode(4)

	l.insertNode(10)
	l.insertNode(20)
	l.insertNode(30)
	l.insertNode(40)
	l.insertNode(50)

	result := mergeTwoLists(&l, &l2)

	display(result)

}
