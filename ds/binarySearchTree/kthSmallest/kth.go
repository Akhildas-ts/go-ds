package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	t := TreeNode{}

	insertBst(&t, 10)
	insertBst(&t, 20)
	insertBst(&t, 30)

	insertBst(&t, 40)
	insertBst(&t, 50)

	fmt.Println(kthSmallest(&t, 2))

}

func insertBst(t *TreeNode, val int) *TreeNode {

	if t == nil {

		return &TreeNode{Val: val, Left: nil, Right: nil}
	}

	if t.Val > val {

		t.Left = insertBst(t.Left, val)

	} else if t.Val < val {

		t.Right = insertBst(t.Right, val)
	}

	return t
}

func kthSmallest(root *TreeNode, k int) int {

	result := inorderTraversal(root)

	return result[k-1]

}

func inorderTraversal(root *TreeNode) []int {

	result := []int{}

	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)

		result = append(result, root.Val)

		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}
