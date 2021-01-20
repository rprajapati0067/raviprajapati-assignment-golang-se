package assignment_1

type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

var (
	result []int
)

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
func (bst *BinarySearchTree) InOrder() []int {
	result = nil
	return bst.InOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) InOrderTraversal(node *TreeNode) []int {

	if node == nil {
		return result
	}
	bst.InOrderTraversal(node.leftNode)
	result = append(result, node.value)
	bst.InOrderTraversal(node.rightNode)
	return result
}

// PreOrder
func (bst *BinarySearchTree) PreOrder() []int {
	result = nil
	return bst.PreOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) PreOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return result
	}
	result = append(result, node.value)
	bst.InOrderTraversal(node.leftNode)
	bst.InOrderTraversal(node.rightNode)
	return result
}

// PostOrder
func (bst *BinarySearchTree) PostOrder() []int {
	result = nil
	return bst.PostOrderTraversal(bst.root)

}

func (bst *BinarySearchTree) PostOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return result
	}
	bst.InOrderTraversal(node.leftNode)
	bst.InOrderTraversal(node.rightNode)
	result = append(result, node.value)
	return result
}
