package main

import "fmt"

const (
	arraysize = 7
)

type bucket struct {
	head *buckedNode
}

type buckedNode struct {
	key   string
	value string
	next  *buckedNode
}

type hashTable struct {
	arr [arraysize]bucket
}

func hash(key string) int {

	sum := 0

	for _, val := range key {

		sum += int(val)
	}

	x := sum % arraysize

	return x

}

func (h *hashTable) Insert(key, value string) {

	index := hash(key)

	if _, ok := h.Search(key); !ok {

		h.arr[index].insertNode(key, value)
	}

}

func (b *bucket) insertNode(key, value string) {
	newNode := &buckedNode{key: key, value: value}
	newNode.next = b.head
	b.head = newNode
}

func (h *hashTable) Search(key string) (string, bool) {

	index := hash(key)

	return h.arr[index].bucketSearch(key)

}

func (h *bucket) bucketSearch(key string) (string, bool) {

	current := h.head

	if h.head == nil {
		return "", false
	}

	if current.key == key {
		return current.value, true

	}

	for current.next != nil {

		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return "", false

}

func main() {

	h := hashTable{}

	h.Insert("akhil", "das")
	h.Insert("arun", "das")
	h.Insert("surra", "ettan")
	h.Insert("thavala", "punda")

	fmt.Println(h.Search("thavala"))
}

// const (
// 	arraysize = 7
// )

// type bucketNode struct {
// 	key   string
// 	value string
// 	next  *bucketNode
// }
// type bucket struct {
// 	head *bucketNode
// }
// type hashTable struct {
// 	arr [arraysize]*bucket
// }

// func (h *hashTable) Insert(key, value string) {

// 	index := hash(key)

// 	if _, ok := h.Search(key); !ok {

// 		h.arr[index].insert(key, value)
// 	} else {

// 		fmt.Println("already exist ")
// 	}

// }

// func (b *bucket) insert(key, value string) {

// 	newNode := &bucketNode{key: key, value: value}

// 	newNode.next = b.head

// 	b.head = newNode
// }

// func (h *hashTable) Search(key string) (string, bool) {

// 	index := hash(key)
// 	return h.arr[index].bucketSearch(key)
// }

// func (h *bucket) bucketSearch(key string) (string, bool) {

// 	current := h.head

// 	if h.head == nil {
// 		return "", false
// 	}

// 	for current != nil {

// 		if current.key == key {
// 			return current.value, true
// 		}

// 		current = current.next

// 	}

// 	return "", false

// }
// func hash(key string) int {

// 	sum := 0
// 	for _, val := range key {

// 		sum = +int(val)
// 	}

// 	x := sum % arraysize

// 	return x

// }

// func Init() *hashTable {

// 	result := &hashTable{}

// 	for i, _ := range result.arr {

// 		result.arr[i] = &bucket{}

// 	}

// 	return result

// }

// func main() {

// 	h := Init()

// 	h.Insert("akhil", "10")
// 	h.Insert("sumisha", "20")
// 	h.Insert("arun", "30")

// 	value, _ := h.Search("sumisha")

// 	fmt.Println(value)

// }
