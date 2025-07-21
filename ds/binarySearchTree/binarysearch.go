package main

import "fmt"

// time complixity

// how many time to take to completed to finish work
// O (log n )

type tree struct {
	val   int
	left  *tree
	right *tree
}

func insertBst(t *tree, val int) *tree {

	if t == nil {

		return &tree{val: val, left: nil, right: nil}
	}

	if t.val > val {

		t.left = insertBst(t.left, val)

	} else if t.val < val {

		t.right = insertBst(t.right, val)
	}

	return t
}

// left ,node ,right
func inorderTraverse(t *tree) {

	if t == nil {
		return
	}

	inorderTraverse(t.left)
	fmt.Println(t.val)

	inorderTraverse(t.right)
}

func main() {

	t := tree{}

	insertBst(&t, 10)
	insertBst(&t, 20)
	insertBst(&t, 40)
	insertBst(&t, 7)

	inorderTraverse(&t)
}
