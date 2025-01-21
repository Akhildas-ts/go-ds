package main

import "fmt"

//implement binary tree
// traverse in inorder

type tree struct {
	data  int
	left  *tree
	right *tree
}

func insertBtree(root *tree, val int) *tree {

	if root == nil {

		return &tree{data: val, left: nil, right: nil}
	}

	if root.left == nil {

		root.left = insertBtree(root.left, val)
	} else {

		root.right = insertBtree(root.right, val)
	}

	return root
}

func inorderTraver(t *tree) {

	if t == nil {
		return
	}

	inorderTraver(t.left)

	fmt.Println(t.data)
	inorderTraver(t.right)
}
func main() {

	t := tree{}

	insertBtree(&t, 10)
	insertBtree(&t, 20)
	insertBtree(&t, 30)
	insertBtree(&t, 40)
	insertBtree(&t, 50)

	inorderTraver(&t)
}
