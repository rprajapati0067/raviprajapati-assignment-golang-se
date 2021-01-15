package assignment_1

import (
	"fmt"
)

type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

// Insert
func (bst *BinarySearchTree) Insert(value int) {
	bst.insertElement(bst.root, value)
}

func (bst *BinarySearchTree) insertElement(node *TreeNode, val int) *TreeNode {
	if bst.root == nil {
		bst.root = &TreeNode{value: val}
		return bst.root
	}
	if node == nil {
		return &TreeNode{value: val}
	}

	if val < node.value {
		node.leftNode = bst.insertElement(node.leftNode, val)
	}
	if val > node.value {
		node.rightNode = bst.insertElement(node.rightNode, val)
	}
	return node
}

// InOrder
func (bst *BinarySearchTree) InOrder() []interface{} {
	return bst.InOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) InOrderTraversal(node *TreeNode) []interface{} {

	var result []interface{}
	if node != nil {

		bst.InOrderTraversal(node.leftNode)
		fmt.Printf("%d \n", node.value)
		bst.InOrderTraversal(node.rightNode)

		result = append(result, node.leftNode)
		result = append(result, node.value)
		result = append(result, node.rightNode)
	}
	return result
}

// PreOrder
func (bst *BinarySearchTree) PreOrder() {
	bst.PreOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) PreOrderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d \n", node.value)
		bst.PreOrderTraversal(node.leftNode)
		bst.PreOrderTraversal(node.rightNode)
	}
}

// PostOrder
func (bst *BinarySearchTree) PostOrder() {
	bst.PostOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) PostOrderTraversal(node *TreeNode) {
	if node != nil {
		bst.PostOrderTraversal(node.leftNode)
		bst.PostOrderTraversal(node.rightNode)
		fmt.Printf("%d \n", node.value)
	}
}
