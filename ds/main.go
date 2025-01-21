package main

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
		l.head.Data = x
		return

	}

	current := l.head

	for current != nil {

		current = current.Next

	}

	current.Data = x

}
func main() {

	l := linkedlist{}

	l.insertion(10)
	l.insertion(20)
	l.insertion(30)
	l.insertion(40)
	l.insertion(50)
	l.insertion(60)

}
