package main

import (
	"fmt"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
Example tree

           3
         /   \
        5     1
       / \  /  \
      6  2  0   8
        / \
       7   4
*/

// level-order by a queue
func levelOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	for len(queue) != 0 {
		// pop from queue
		head := queue[0]
		*res = append(*res, head.Val)
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
}

// non-recursive pre-order traverse
// access element when put into stack
func nonRecursivePreOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	stack := make([]*TreeNode, 0)

	// pre-step, put root and all its left node into stack
	for node != nil {
		stack = append(stack, node)
		*res = append(*res, node.Val)
		node = node.Left
	}

	for len(stack) > 0 {
		// step 1, pop one element from stack
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		// step 2, access right node
		if top.Right != nil {
			stack = append(stack, top.Right)
			*res = append(*res, top.Right.Val)

			// step 3, put right node and all its left node into stack
			left := top.Right.Left
			for left != nil {
				stack = append(stack, left)
				*res = append(*res, left.Val)
				left = left.Left
			}
		}
	}
}

// access element when pop form stack
func nonRecursiveInOrder(node *TreeNode, res *[]int) {

	stack := make([]*TreeNode, 0)
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		*res = append(*res, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
			left := top.Right.Left

			for left != nil {
				stack = append(stack, left)
				left = left.Left
			}
		}
	}
}

// non-recursive post order, is it suitable to use a set ?
// when assess a nodes right child, put it in a set
func nonRecursivePostOrder(node *TreeNode, res *[]int) {
	stack := make([]*TreeNode, 0)
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	sets := make(map[*TreeNode]struct{})
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		_, exist := sets[top]

		// the top element has no right child, or already been accessed, pop
		if top.Right == nil || exist {
			// pop
			stack = stack[0 : len(stack)-1]
			*res = append(*res, top.Val)
			continue
		}

		// the top element has right child, do not pop, access its right child
		if top.Right != nil {
			// if right child is not nil, do not pop out of stack, access its right child
			sets[top] = struct{}{}
			right := top.Right
			for right != nil {
				stack = append(stack, right)
				right = right.Left
			}
		}
	}
}

// recursive pre-order traverse
func preOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	preOrder(node.Left, res)
	preOrder(node.Right, res)
	return
}

// recursive in-order traverse
func inOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, res)
	*res = append(*res, node.Val)
	inOrder(node.Right, res)
}

// recursive post-order traverse
func postOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, res)
	postOrder(node.Right, res)
	*res = append(*res, node.Val)
}

func main() {
	root := ConstructExampleTree()

	res := make([]int, 0)
	//preOrder(root, &res)
	nonRecursivePostOrder(root, &res)
	fmt.Println(res)
}

func ConstructExampleTree() *TreeNode {
	root := &TreeNode{Val: 3}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node1 := &TreeNode{Val: 1}

	root.Left = node5
	root.Right = node1
	node5.Left = node6
	node5.Right = node2
	node2.Left = node7
	node2.Right = node4
	node1.Left = node0
	node1.Right = node8

	return root
}
